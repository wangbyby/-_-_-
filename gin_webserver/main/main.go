package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"repro_go/gin_webserver/http/controller"
	"time"

	"github.com/gin-gonic/gin"
)

//Loggeer中间件
func Logger() gin.HandlerFunc { //感觉现在没啥用
	return func(c *gin.Context) {
		t := time.Now()
		// 设置 example 变量
		c.Set("example", "12345")
		// 请求前
		c.Next()
		// 请求后
		latency := time.Since(t)
		log.Print(latency)
		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}

//日志写入文件+控制台
// gin.DisableConsoleColor()
// f, _ := os.Create("/gin.log")
// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
func main() {

	//新建中间件
	r := gin.New() //
	r.Use(Logger())
	// r := gin.Default() //使用默认中间件

	r.LoadHTMLGlob("../view/*") //载入html
	//静态文件
	r.StaticFile("/favicon.png", "./static/zu.png")
	r.MaxMultipartMemory = 8 << 21 //上传文件最大 80MB
	frontGroup := r.Group("")      //分组注册
	controller.RegisterRoutes(frontGroup)

	fmt.Println("listen on: localhost:8000")
	//关闭服务器
	srv := &http.Server{
		Addr:    ":8000", //监听8000端口
		Handler: r,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	//关闭所有句柄
	defer cancel()
	defer controller.AntsPool.Release() //线程池
	defer controller.RedisPtr.Close()   // redis
	defer controller.PgDB.Close()       //pgsql

	//关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	//r.Run() // listen and serve on 0.0.0.0:8000 (8080被tomcat占了)

}
