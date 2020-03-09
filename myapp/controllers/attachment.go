package controllers

import (
	"io"
	"net/url"
	"os"

	"github.com/astaxie/beego"
)

type AttachController struct {
	beego.Controller
}

func (this *AttachController) Get() {
	filePath, err := url.QueryUnescape(this.Ctx.Request.RequestURI[1:])
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	_, err = io.Copy(this.Ctx.ResponseWriter, f)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
}
