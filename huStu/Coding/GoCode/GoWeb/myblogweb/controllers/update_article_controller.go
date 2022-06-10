package controllers

import (
	"fmt"
	"myblogweb/models"
	"github.com/astaxie/beego"
)


type UpdateArticleController struct {
	BaseController
}


func (this *UpdateArticleController) Get() {
	id,_ := this.GetInt("id")
	fmt.Println("update article id -----------> ",id)

	//获取id对应的文章信息
	art,_ := models.QueryArticleWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Tags"] = art.Tags
	this.Data["Short"] = art.Short
	this.Data["Content"] = art.Content
	this.Data["Id"] = art.Id
	this.TplName = "write_article.html"
}


func (this *UpdateArticleController) Post() {
	id,_ := this.GetInt("id")
	fmt.Println("postid : ",id)

	//获取浏览器传输的数据
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	//实例化一个Article
	art := models.Article{id,title,"",tags,short,content,0}
	_,err := models.UpdateArticle(art)

	//返回数据给浏览器
	if err!=nil {
		this.Data["json"] = map[string]interface{}{"code":0,"message":"更新失败"}
	} else {
		this.Redirect(beego.URLFor("HomeController.Get"),302)
		this.Data["json"] = map[string]interface{}{"code":1,"message":"更新成功"}
	}
	this.ServeJSON()
}