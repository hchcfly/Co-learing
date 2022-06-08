package controllers

import (
	"fmt"
	"myblogweb/models"
)


type HomeController struct {
	BaseController
}


func (this *HomeController) Get() {
	//  获取分页查询页数
	page,_ := this.GetInt("page")
	fmt.Println("page: ",page)
	if page <= 0{
		page = 1
	}
	var artList []models.Article
	artList , _ = models.FindArticleWithPage(page)
	this.Data["pageCode"] = 1
	this.Data["HasFooter"] = true

	fmt.Println("IsLogin:",this.IsLogin,this.Loginuser)
	this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)
	this.TplName = "home.html"
}