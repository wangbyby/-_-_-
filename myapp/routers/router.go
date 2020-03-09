package routers

import (
	"myapp/controllers"
	"os"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	beego.Router("/topic", &controllers.TopicController{})

	beego.AutoRouter(&controllers.TopicController{})

	//附件处理
	os.MkdirAll("attachment", os.ModePerm)
	//1.作为静态文件
	// beego.SetStaticPath("/attachment", "attachment")
	//2.作为控制器
	beego.Router("/attachment/:all", &controllers.AttachController{})
}
