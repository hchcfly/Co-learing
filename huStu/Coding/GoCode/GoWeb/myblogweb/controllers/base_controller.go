package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)


type BaseController struct {
	beego.Controller
	IsLogin bool  //  用户是否登录
	Loginuser interface{}  //表示用户名
}


//判断是否登录
//该函数会在调用其他方法前自动调用
func (this *BaseController) Prepare() {
    loginuser := this.GetSession("loginuser")
    fmt.Println("loginuser---->", loginuser)
    if loginuser != nil {
        this.IsLogin = true
        this.Loginuser = loginuser
    } else {
        this.IsLogin = false
    }
    this.Data["IsLogin"] = this.IsLogin
}





