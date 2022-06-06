package routers

import (
	"myblogweb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/register", &controllers.RegisterController{})  //注册路由

	beego.Router("/login", &controllers.LoginController{})   //登录路由 
}
