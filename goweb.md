go web
个人博客...
1. 项目设计
   1. 结构 : MVC
   2. web框架 : beego
   3. 探究web底层实现
      1. net/http包, http.HandleFunc
        ```go
        package main
        import (
            "io"
            "log"
            "net/http"
        )
        func main() {
            http.HandleFunc("/", handler)
            err := http.ListenAndServe(":8088", nil)
            if err != nil {
                log.Fatal(err)
            }
        }
        func handler(w http.ResponseWriter, r *http.Request) {
            io.WriteString(w, "hello version 1 ")
        }
        ```
        2.  进一步探究web
        ```go
        package main

        import (
            "io"
            "log"
            "net/http"
        )

        func main() {
            mux := http.NewServeMux()
            mux.Handle("/", &myHandler{})
            mux.HandleFunc("/hello", sayHello)

            //静态文件的实现
            wd, err1 := os.Getwd()
            if err1 != nil {
                log.Fatal(err1)
            }
            mux.Handle("/static/", http.StripPrefix("/static/",
                http.FileServer(http.Dir(wd))))
                    err := http.ListenAndServe(":8088", mux)
            if err != nil {
                log.Fatal(err)
            }
        }

        type myHandler struct{}

        func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
            io.WriteString(w, "hello version 2 URL: "+r.URL.String())
        }

        func sayHello(w http.ResponseWriter, r *http.Request) {
            io.WriteString(w, "hello version 2 ")
        }

        
        ```
        3. web底层实现, ServeHTTP
        ```go
        package main

        import (
            "io"
            "log"
            "net/http"
            "os"
            "time"
        )

        var mux map[string]func(http.ResponseWriter, *http.Request)

        func main() {

            server := http.Server{
                Addr:              ":8080",
                Handler:           &myHandler{},
                ReadHeaderTimeout: 5 * time.Second,
            }


            mux = make(map[string]func(http.ResponseWriter, *http.Request))
            mux["/"] = sayHello
            mux["/bye"] = sayBye

            err = server.ListenAndServe()
            if err != nil {
                log.Fatal(err)
            }
        }

        type myHandler struct{}

        func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
            //路由转发
            if h, ok := mux[r.URL.String()]; ok {
                h(w, r)
                return
            }
            io.WriteString(w, "hello version 3 URL: "+r.URL.String())
        }

        func sayHello(w http.ResponseWriter, r *http.Request) {
            io.WriteString(w, "hello version 3 ")
        }

        func sayBye(w http.ResponseWriter, r *http.Request) {
            io.WriteString(w, "88 version 3")
        }

        
        ```
   4. beego实际开发
      1. 模板用法讲解 
        ```go
        
        func (c *MainController) Get() {

            c.Data["Website"] = "beego.me"
            c.Data["Email"] = "astaxie@gmail.com"
            c.TplName = "index.tpl"

            c.Data["TrueCond"] = true
            c.Data["FalseCond"] = false

            type u struct {
                Name string
                Age  int
                Sex  string
            }
            user := &u{
                Name: "wang",
                Age:  10,
                Sex:  "男",
            }
            c.Data["User"] = user

            nums := []int{1, 2, 3, 4, 1, 3, 1, 10}
            c.Data["nums"] = nums

            c.Data["tplVar"] = "hello world" //模板变量

            //模板函数
            c.Data["html"] = "<div>hello bingo</div>"

            c.Data["pipe"] = "<div>hello bangbang</div>"
        }

        ```
      2. ORM设计
        ```go
        package models
        import (
            "os"
            "path"
            "time"
            "github.com/Unknwon/com"
            "github.com/astaxie/beego/orm"
            _ "github.com/lib/pq"
        )
        const (
            _DB_NAME   = "host=localhost port=5432 user=postgres password=root dbname=default sslmode=disable"
            _PG_DRIVER = "postgres"
        )
        //分类
        type Category struct {
            Id              int64
            Title           string
            Created         time.Time `orm:"index"`
            Views           int64     `orm:"index"`
            TopicTime       time.Time `orm:"index"`
            TopicCount      int64
            TopicLastUserId int64
        }
        //文章
        type Topic struct {
            Id              int64
            Uid             int64
            Title           string
            Content         string `orm:"size(5000)"`
            Attachment      string
            Created         time.Time `orm:"index"`
            Updated         time.Time `orm:"index"`
            Views           int64     `orm:"index"`
            Author          string
            Replytime       time.Time `orm:"index"`
            ReplyCount      int64
            ReplyLastUserId int64
        }
        func RegisterDB() {
            if !com.IsExist(_DB_NAME) {
                os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
                os.Create(_DB_NAME)
            }
            orm.RegisterModel(new(Category), new(Topic))
            orm.RegisterDriver(_PG_DRIVER, orm.DRPostgres)
            orm.RegisterDataBase("default", _PG_DRIVER, _DB_NAME, 10)
        }
        ```
- go-rpc
  - http 
  - tcp
  - json
  - 代码
    1. client
        ```go
        package main
        import (
            "fmt"
            "log"

            "net/rpc/jsonrpc"
            "os"
        )

        type Args struct {
            A, B int
        }

        type Math int

        type Result struct {
            Quo, Rem int
        }

        func main() {

            if len(os.Args) != 2 {
                fmt.Println("usage : ", os.Args[0], " server")
                os.Exit(1)
            }
            serverAdd := os.Args[1]
            //DialHTTP
            //Dial
            client, err := jsonrpc.Dial("tcp", serverAdd+":8088")
            defer client.Close()
            if err != nil {
                log.Println("tcp dial error with", err)
            }

            args := Args{17, 19}
            var reply int
            err = client.Call("Math.Mul", args, &reply)
            if err != nil {
                log.Println("rpc call error with", err)
            }
            fmt.Printf("\t%v*%v=%v\n", args.A, args.B, reply)

            var quo Result
            err = client.Call("Math.Div", args, &quo)
            if err != nil {
                log.Println("rpc call error with", err)
            }
            fmt.Printf("\t%v/%v=%v...%v\n", args.A, args.B, quo.Quo, quo.Rem)
        }
        ```
    2. server
        ```go
        package main
        import (
            "errors"
            "fmt"
            "net"
            "net/rpc"
            "net/rpc/jsonrpc"
            "os"
        )
        func main() {

            m := new(Math)
            rpc.Register(m)
            // rpc.HandleHTTP()
            tcpAdd, err := net.ResolveTCPAddr("tcp", ":8088")
            if err != nil {
                fmt.Println("TCP error", err)
                os.Exit(2)
            }
            listener, err := net.ListenTCP("tcp", tcpAdd)
            if err != nil {
                fmt.Println("listen TCP error", err)
                os.Exit(2)
            }

            for {
                conn, err := listener.Accept()
                if err != nil {
                    fmt.Println("accept TCP error", err)
                    continue
                }
                // rpc.ServeConn(conn)
                jsonrpc.ServeConn(conn)
            }

        }

        type Args struct {
            A, B int
        }

        type Math int

        func (m *Math) Mul(args *Args, reply *int) error {
            *reply = args.A * args.B
            return nil
        }

        type Result struct {
            Quo, Rem int
        }

        func (m *Math) Div(args *Args, reply *Result) error {
            if args.B == 0 {
                return errors.New("div by zero")
            }
            reply.Quo = args.A / args.B
            reply.Rem = args.A % args.B
            return nil
        }
        ```
- REST & websocket
    
