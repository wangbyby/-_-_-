package controller

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"repro_go/gin_webserver/dao"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/panjf2000/ants"
	"github.com/redigo-master/redis"
)

var RedisPtr redis.Conn //redis句柄
var PgDB *sql.DB        //pgsql句柄
var AntsPool *ants.Pool //线程池句柄
var errg error          //全局error 只用于init中初始化异常

// var admin map[string]int

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
	// admin = make(map[string]int, 5)
}

type IndexController struct{}

func (i IndexController) RegisterRoute(r *gin.RouterGroup) {
	//http重定向
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})
	//路由重定向
	r.GET("/test1", func(c *gin.Context) {
		//重定向到根目录
		redirctPath(c, "/")
	})

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

	//上传文件
	//单个文件
	r.POST("/uploadone", func(c *gin.Context) {
		uploadOneFile(c)
	})
	//多个文件
	r.POST("/upload", func(c *gin.Context) {
		/* 使用副本,会报错... uploadFiles(c.Copy()) */
		uploadFiles(c)
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

/*
--例如
	查询 /query?type=query&id=1
	删除 /query?type=delete&name=烤鸭
	新增 /query?type=insert&name=烤鸭&price=100&class=肉菜&imgurl=1
	更改 /query?type=update&oldname=烤鸭&newname=北京烤鸭&price=300&class=肉菜
*/
func queryGET(c *gin.Context) {
	switch c.Query("type") {
	case "query":
		{
			info := gin.H{}
			syncChan := make(chan int)
			if err := ants.Submit(func() {
				queryGETqueryPool(c, syncChan, info)
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

func queryGETqueryPool(c *gin.Context, syncChan chan int, dishInfo gin.H) {
	var idq string
	var name string
	var class string
	var price []uint8
	var imgurl string
	id := c.Query("id")
	if idint, _ := strconv.Atoi(id); idint == 0 { // id为0时, 查找全部
		sqlcontent := `select "id", "name","class","price","imgurl" from public.dish`
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
			row.Scan(&idq, &name, &class, &price, &imgurl)
			tmp := make(map[string]string, 4)
			tmp["id"] = idq
			tmp["菜系"] = class
			tmp["菜名"] = name
			tmp["价格"] = string(price)
			tmp["imgurl"] = imgurl
			dishInfo[strconv.Itoa(i)] = tmp
			i++
		}
	} else {
		sqlcontent := `select "id", "name","class","price", "imgurl" from public.dish where id=$1`
		row := PgDB.QueryRow(sqlcontent, id) //只查一行

		err := row.Scan(&idq, &name, &class, &price, &imgurl)
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
		dishInfo["imgurl"] = imgurl
	}
	syncChan <- 1
}

//属于 quertGET的子函数 处理type = insert
//Ex.	/query?type=insert&name=宫保鸡丁&class=肉菜&price=20&imgurl=10
func queryGETinsert(c *gin.Context) {

	imgurl := c.Query("imgurl")

	name := c.Query("name")
	class := c.Query("class")
	price := c.Query("price")

	sqlcontent := `insert into public.dish (name,class,price,imgurl) values ($1,$2,$3,$4)`
	stmt, err := PgDB.Prepare(sqlcontent)
	_, err = stmt.Exec(name, class, price, imgurl)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"flag": 0})
		return
	}
	c.JSON(http.StatusOK, gin.H{"flag": 1})
}

//属于 quertGET的子函数 处理type = delete
// /query?type=delete&name=宫保鸡丁
func queryGETdelete(c *gin.Context) {

	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "bad request", "flag": 0})
		return
	}

	sqlDel := `delete from public.dish where name=$1`
	stmt, err := PgDB.Prepare(sqlDel)
	_, err = stmt.Exec(name)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"flag": 0}) //删除失败
		return
	}
	c.JSON(http.StatusOK, gin.H{"flag": 1}) //删除成功
}

//属于 quertGET的子函数 处理type = update
//Ex. /query?type=update&oldname=烤鸭&class=肉菜&price=20&newname=北京烤鸭
func queryGETupdate(c *gin.Context) {

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
	var name, pwd string
	var ulevel int
	//获取用户信息
	if err := c.ShouldBind(&form); err != nil {
		/*
			c.ShouldBindBodyWith 会在绑定之前将 body 存储到上下文中。 这会对性能造成轻微影响，如果调用一次就能完成绑定的话，那就不要用这个方法
		*/
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	online, _ := UserOnline(form.User)
	sqlQuery := `select  "uname","upwd","ulevel" from public.usert where uname=$1`
	row := PgDB.QueryRow(sqlQuery, form.User)
	err := row.Scan(&name, &pwd, &ulevel)
	if !online { //redis中没有该用户
		/*
			1.如果数据库中存在该用户
				则登录,添加到redis中去
			2.如果数据库中不存在该用户
				添加到数据库, redis中去
		*/
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
				"ulevel":  0,
			})
			return
		}
		//1.如果数据库中存在该用户
		//则比较密码
		if pwd != form.Password { //密码错误
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error password"})
			return
		}
		//密码正确
		//登录,添加到redis中去
		RedisPtr.Do("SADD", "sessionID", form.User)
		//返回主页面
		loginPostTaskReturn(c, ulevel)
		return
	}
	// redis 中存在该用户
	//返回主页面

	if pwd != form.Password || err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error password"})
		return
	}
	loginPostTaskReturn(c, ulevel)
}

func loginPostTaskReturn(c *gin.Context, ulevel int) {
	if ulevel == 0 {
		c.HTML(http.StatusOK, "../view/root.htm", gin.H{
			"doc":     "点菜系统",
			"logined": true,
			"ulevel":  ulevel,
		})
		return
	}
	c.HTML(http.StatusOK, "../view/admin.htm", gin.H{
		"doc":     "管理系统",
		"logined": true,
		"ulevel":  ulevel,
	})
}

//路由重定向
func redirctPath(c *gin.Context, routerStr string) {
	c.Redirect(http.StatusMovedPermanently, routerStr)
}

func getFileId(dstPath string) (int, error) {

	tmp, err := ioutil.ReadDir(dstPath)
	if err != nil {

		return 0, err
	}
	dirLen := len(tmp) //得到文件夹的文件数量
	dirLen++
	return dirLen, nil
}

//上传单文件...
func uploadOneFile(c *gin.Context) {
	//优化 : 以时间和序列号作为文件名..
	file, _ := c.FormFile("file")
	dst := file.Filename
	log.Println(dst)

	// 上传文件到指定的路径
	c.SaveUploadedFile(file, "../static/images/"+dst)
	c.JSON(http.StatusOK, gin.H{"flag": 1, "url": dst})

}

//上传多文件处理
func uploadFiles(c *gin.Context) {
	dstPath := "../static/images/"
	dirLen, _ := getFileId(dstPath)
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	fileNameInfo := gin.H{}
	for _, file := range files {
		tmp := strconv.Itoa(dirLen)
		fileName := dstPath + tmp + ".jpg"
		dirLen++
		fileNameInfo[tmp] = tmp
		//写入文件

		if err := c.SaveUploadedFile(file, fileName); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"flag": 0})
			return
		}
	}
	fileNameInfo["flag"] = 1
	c.JSON(http.StatusOK, fileNameInfo)
}
