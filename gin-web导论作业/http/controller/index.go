package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"repro_go/gin_webserver/dao"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/panjf2000/ants"
	"github.com/redigo-master/redis"
)

var RedisPtr redis.Conn //redis句柄
var PgDB *sql.DB        //pgsql句柄
var AntsPool *ants.Pool //线程池句柄
var errg error          //全局error 只用于init中初始化异常

func init() {
	//若不能打开redis以及pgsql,以及初始化线程池错误则直接退出
	//打开redis
	RedisPtr, errg = redis.Dial("tcp", "127.0.0.1:6379") //向本机链接
	if errg != nil {
		log.Println("dial redis error\n", errg)
		os.Exit(0)
	}

	//打开pgsql 用户是postgres 密码 root
	PgDB, errg = sql.Open("postgres", "host=localhost port=5432 user=postgres password=root dbname=postgres sslmode=disable")
	if errg != nil {
		log.Println("dial pgsql error\n", errg)
		os.Exit(0)
	}
	//初始化线程池
	max := 1000000
	AntsPool, errg = ants.NewPool(max)
	if errg != nil {
		log.Println("init antsPool error", errg)
		os.Exit(0)
	}
}

type IndexController struct{}

func (i IndexController) RegisterRoute(r *gin.RouterGroup) {
	// //路由重定向
	// r.GET("/test1", func(c *gin.Context) {
	// 	c.Request.URL.Path = "/test2"
	// 	// r.HandleContext(c)
	// })
	// r.GET("/test2", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"hello": "world"})
	// })

	//设置 获取 Cookie
	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value is : %v\n", cookie)
	})
	//GET 匹配 Ex. /user/jojo/ola
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)

	})
	//GET / (根目录)
	r.GET("/", root)

	//上传多文件
	r.POST("/upload", func(c *gin.Context) {
		//并发 必须使用副本
		err := AntsPool.Submit(func() { uploadFiles(c.Copy()) })
		if err != nil {
			log.Println("上传文件出错", err)
		}
	})
	//下载文件
	r.GET("/download", downloadFile)

	//查询数据库
	r.GET("/query", queryGET)

	//判断用户是否登录
	r.GET("/logined", loginedGET)
	//用户登录
	r.POST("/login", loginPOST)
	//加载login的html文件
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "../view/login.htm", gin.H{
			"title": "登录",
		})
	})

	//http重定向
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

}

/*以下是具体实现的函数*/

func loginedGET(c *gin.Context) {
	var form dao.Login
	//获取用户信息
	if err := c.ShouldBind(&form); err != nil {
		/*
			c.ShouldBindBodyWith 会在绑定之前将 body 存储到上下文中。 这会对性能造成轻微影响，如果调用一次就能完成绑定的话，那就不要用这个方法
		*/
		c.JSON(http.StatusOK, gin.H{"logined": 0})
		return
	}
	online, _ := UserOnline(form.User)
	if online {
		c.JSON(http.StatusOK, gin.H{"logined": 1})
	} else {
		c.JSON(http.StatusOK, gin.H{"logined": 0})
	}
}

// /download?url=a.jpg
func downloadFile(c *gin.Context) {
	//根据地址下载文件
	url := c.Query("url")
	c.File("../static/images/" + url)
}

//UserOnline... 用户是否在redis缓存中
func UserOnline(username string) (online bool, err error) {
	id, err := RedisPtr.Do("SISMEMBER", "sessionID", username)
	if id.(int64) == 0 { //用户不在线
		return false, err
	}
	return true, err
}

// 根目录
func root(c *gin.Context) {
	c.HTML(http.StatusOK, "../view/root.htm", gin.H{
		"doc": "点菜系统",
	})
}

//数据库增删改查
//匹配方式: /query?type=[query,delete,insert,update]
//例如
/*
	查询 /query?type=query&id=1
	删除 /query?type=delete&id=1
	新增 /query?type=insert&name=烤鸭&price=100&class=鲁菜
	更改 /query?type=update&oldname=烤鸭&newname=北京烤鸭&price=300&class=荤菜
*/
func queryGET(c *gin.Context) {
	switch c.Query("type") {
	case "query":
		{
			info := gin.H{}
			syncChan := make(chan int)
			if err := ants.Submit(func() {
				poolqueryGETquery(c, syncChan, info)
			}); err != nil {
				log.Println(err)
			}
			<-syncChan
			c.JSON(http.StatusOK, info)
		}
	case "delete":
		{
			queryGETdelete(c)
		}
	case "insert":
		{
			queryGETinsert(c)
		}
	case "update":
		{
			queryGETupdate(c)
		}
	default:
		{
			c.HTML(http.StatusOK, "view/query.htm", gin.H{
				"doc": "Hello gin",
			})
		}
	}
}

/*
--查询 属于 quertGET的子函数 处理type = query
--url : /query?type=query&id=1
--当id==0时, 返回所有结果
*/

func poolqueryGETquery(c *gin.Context, syncChan chan int, dishInfo gin.H) {
	var idq string
	var name string
	var class string
	var price []uint8
	id := c.Query("id")
	if idint, _ := strconv.Atoi(id); idint == 0 {
		sqlcontent := `select "id", "name","class","price" from public.dish`
		row, err := PgDB.Query(sqlcontent)
		if err != nil {
			log.Println(err)
			dishInfo["err"] = err.Error()
			syncChan <- 0
			return
		}
		i := 1
		dishInfo["flag"] = "all"
		for row.Next() {
			row.Scan(&idq, &name, &class, &price)
			tmp := make(map[string]string, 4)
			tmp["id"] = idq
			tmp["菜系"] = class
			tmp["菜名"] = name
			tmp["价格"] = string(price)
			dishInfo[strconv.Itoa(i)] = tmp
			i++
		}
	} else {
		sqlcontent := `select "id", "name","class","price" from public.dish where id=` + id
		row := PgDB.QueryRow(sqlcontent) //只查一行

		err := row.Scan(&idq, &name, &class, &price)
		if err != nil {
			log.Println(err)
			dishInfo["err"] = err.Error()
			syncChan <- 0
			return
		}
		dishInfo["flag"] = "one"
		dishInfo["id"] = idq
		dishInfo["菜系"] = class
		dishInfo["菜名"] = name
		dishInfo["价格"] = string(price)
	}
	syncChan <- 1
}

/*
func queryGETquery(c *gin.Context) {

	dishInfo := make(map[string]interface{})
	var idq string
	var name string
	var class string
	var price []uint8
	id := c.Query("id")
	if idint, _ := strconv.Atoi(id); idint == 0 {
		sqlcontent := `select "id", "name","class","price" from public.dish`
		row, err := PgDB.Query(sqlcontent)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"err": "query error"})
		}
		i := 0
		dishInfo["flag"] = "all"
		for row.Next() {
			row.Scan(&idq, &name, &class, &price)
			tmp := make(map[string]string, 4)
			tmp["id"] = idq
			tmp["菜系"] = class
			tmp["菜名"] = name
			tmp["价格"] = string(price)
			dishInfo[string(i)] = tmp
			i++
		}
	} else {
		sqlcontent := `select "id", "name","class","price" from public.dish where id=` + id
		row := PgDB.QueryRow(sqlcontent) //只查一行

		err := row.Scan(&idq, &name, &class, &price)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"err": "query error"})
		}
		dishInfo["flag"] = "one"
		dishInfo["id"] = idq
		dishInfo["菜系"] = class
		dishInfo["菜名"] = name
		dishInfo["价格"] = string(price)
	}
	c.JSON(http.StatusOK, dishInfo)
}
*/
//属于 quertGET的子函数 处理type = insert
func queryGETinsert(c *gin.Context) {
	//Ex.	/query?type=insert&name=宫保鸡丁&class=鲁菜&price=20
	name := c.Query("name")
	class := c.Query("class")
	price := c.Query("price")
	sqlcontent := `insert into public.dish (name,class,price) values ($1,$2,$3)`
	stmt, err := PgDB.Prepare(sqlcontent)
	_, err = stmt.Exec(name, class, price)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"err": "insert error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": "no error"})
}

//属于 quertGET的子函数 处理type = delete
func queryGETdelete(c *gin.Context) {
	// /query?type=delete&id=1
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "bad request"})
		return
	}

	sqlDel := `delete from public.dish where id=$1`
	stmt, err := PgDB.Prepare(sqlDel)
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"err": "delete error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": "no error"})
}

//属于 quertGET的子函数 处理type = update
func queryGETupdate(c *gin.Context) {
	//  /query?type=update&oldname=烤鸭&class=鲁菜&price=20&newname=北京烤鸭
	oldname := c.Query("oldname")
	class := c.Query("class")
	price := c.Query("price")
	newname := c.Query("newname")
	sqlupdate := `update public.dish set name=$1,class=$2,price=$3 where name=$4`
	stmt, err := PgDB.Prepare(sqlupdate)
	_, err = stmt.Exec(newname, class, price, oldname)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"err": "update error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": "no error"})
}

//用户登录
func loginPOST(c *gin.Context) {
	loginPostTask(c)
}

func loginPostTask(c *gin.Context) {
	var form dao.Login
	//获取用户信息
	if err := c.ShouldBind(&form); err != nil {
		/*
			c.ShouldBindBodyWith 会在绑定之前将 body 存储到上下文中。 这会对性能造成轻微影响，如果调用一次就能完成绑定的话，那就不要用这个方法
		*/
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	online, _ := UserOnline(form.User)
	if !online { //redis中没有该用户
		/*
			1.如果数据库中存在该用户
				则登录,添加到redis中去
			2.如果数据库中不存在该用户
				添加到数据库, redis中去
		*/
		var name string
		sqlQuery := `select  "uname" from public.usert where uname=` + form.User
		row := PgDB.QueryRow(sqlQuery)
		err := row.Scan(&name)
		if name == "" && err != nil {
			log.Println("pgsql", err)
			//2.数据库中不存在该用户
			//添加到数据库, redis中去
			sqlQuery = `insert into public.usert (uname,upwd) values ($1,$2)`
			stmt, err := PgDB.Prepare(sqlQuery)
			_, err = stmt.Exec(form.User, form.Password)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusUnauthorized, gin.H{"status": "sorry can not register"})
				return
			}
			//加入redis
			RedisPtr.Do("SADD", "sessionID", form.User)
			c.HTML(http.StatusOK, "../view/root.htm", gin.H{
				"doc":     "点菜系统",
				"logined": true,
			})
			return
		}
		//1.如果数据库中存在该用户
		//则登录,添加到redis中去
		RedisPtr.Do("SADD", "sessionID", form.User)
		//返回主页面
	}
	// redis 中存在该用户
	//返回主页面

	c.HTML(http.StatusOK, "../view/root.htm", gin.H{
		"doc":     "点菜系统",
		"logined": true,
	})
}

func redirctPath(c *gin.Context, routerStr string) {
	host := strings.Split(c.Request.Host, ":")
	host[1] = "8000"
	target := "http://" + strings.Join(host, ":") + routerStr
	c.Redirect(http.StatusMovedPermanently, target)
}

//上传文件处理
func uploadFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	for _, file := range files {
		log.Println(file.Filename)

		//写入文件
		fileName := "D:\\#test/" + file.Filename //配置路径
		if err := c.SaveUploadedFile(file, fileName); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			log.Println(err)
			return
		}
	}
	c.String(http.StatusOK, "%d files uploaded!", len(files))
}
