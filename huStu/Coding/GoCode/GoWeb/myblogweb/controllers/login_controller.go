package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"myblogweb/models"
	"myblogweb/utils"
)

type LoginController struct {
	beego.Controller
}


func (this *LoginController) Get() {
	this.TplName = "login.html"
}


func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:",username,"password:",password)
	password = utils.MD5(password)
	fmt.Println("md5后：",password)
	id := models.QueryUserWithParam(username,password)
	if id > 0{
		this.SetSes sion("loginuser",username)
		this.Redirect(beego.URLFor("HomeController.Get"),302)
		this.Data["json"] = map[string]interface{}{"code":1,"message":"登录成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code":0,"message":"登录失败"}
	}
	this.ServeJSON()
}


