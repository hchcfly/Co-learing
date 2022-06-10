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

	beego.Router("/article/:id",&controllers.ShowArticleController{})  //显示文章内容

	beego.Router("/article/update",&controllers.UpdateArticleController{}) //更新文章路由

	beego.Router("/article/delete",&controllers.DeleteArticleController{}) //删除文章路由

	beego.Router("/tags",&controllers.TagsController{}) //删除文章路由
	
	//相册
	beego.Router("/album", &controllers.AlbumController{})

	//文件上传
	beego.Router("/upload", &controllers.UploadController{})

	//关于我
	beego.Router("/aboutme",&controllers.AboutMeController{})

}
