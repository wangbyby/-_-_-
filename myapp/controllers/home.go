package controllers

import (
	"myapp/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "home.html"
	this.Data["IsHome"] = true

	this.Data["IsLogin"] = checkAccount(this.Ctx)
	topics, err := models.GetAllTopic(this.Input().Get("cate"), this.Input().Get("label"), true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

	categories, err := models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}

	this.Data["Categories"] = categories
}
