package routers

import (
	"myblogweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	beego.Router("/register", &controllers.RegisterController{})  //注册路由

	beego.Router("/login", &controllers.LoginController{})   //登录路由 

	beego.Router("/exit",&controllers.ExitController{})  //退出路由

	beego.Router("/article/add",&controllers.AddArticleController{})  //写文章路由

	//beego.Router("/article/:id",&controllers.ShowArticleController{})  //
}
