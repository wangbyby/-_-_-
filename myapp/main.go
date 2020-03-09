package main

import (
	"myapp/models"
	_ "myapp/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true) //自建表
	
	beego.Run()
}
