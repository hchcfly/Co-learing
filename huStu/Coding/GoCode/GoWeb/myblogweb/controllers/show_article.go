package controllers

import (
	"fmt"
	"myblogweb/models"
	"strconv"
)


type ShowArticleController struct {
	BaseController
}

func (this *ShowArticleController) Get() {
	idStr := this.Ctx.Input.Param(":id")
	id,_ := strconv.Atoi(idStr)
	fmt.Println("id------------>",id)

	//获取id对应的文件信息
	art,_ := models.QueryArticleWithId(id)

	this.Data["Title"] = art.Title
	this.Data["Content"] = art.Content
	this.TplName = "show_article.html"
}