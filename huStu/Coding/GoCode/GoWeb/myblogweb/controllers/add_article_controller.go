package controllers

import (
	"fmt"
	"time"
	"myblogweb/models"
	"github.com/astaxie/beego"
)

type AddArticleController struct {
	BaseController
}


//  GET请求对应的处理
func (this *AddArticleController) Get() {
	this.TplName = "write_article.html"
}

func (this * AddArticleController) Post() {
	//  获取浏览器传输的数据
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	fmt.Printf("title: %s , tags: %s\n",title,tags)

	//  实例化一个Article对象
	art := models.Article{0,title,"hchc",tags,short,content,time.Now().Unix()}
	//  文章数据插入数据库中
	_ ,err := models.AddArticle(art)

	//  返回给浏览器的数据
	var response map[string]interface{}
	if err == nil {
		//  登录成功页面跳转
		this.Redirect(beego.URLFor("HomeController.Get"),302)
		response = map[string]interface{}{"code":1,"message":"ok"}
	} else {
		response = map[string]interface{}{"code":0,"message":"error"}
	}
	this.Data["json"] = response
	this.ServeJSON()

}