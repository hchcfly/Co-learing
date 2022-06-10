package controllers

import (
	"fmt"
	"myblogweb/models"
)


type HomeController struct {
	BaseController
}


func (this *HomeController) Get() {
	//  获取标签值
	tag := this.GetString("tag")
	//  获取分页查询页数
	page,_ := this.GetInt("page")
	var artList []models.Article

	//先接收tag的值再接收page的值
	if len(tag) > 0{
		//按照指定的标签进行搜索
		artList,_ = models.QueryArticleWithTag(tag)
		this.Data["HasFooter"] = false
	} else
	{
		if page <= 0{
			page = 1
		}
		artList , _ = models.FindArticleWithPage(page)
		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		this.Data["HasFooter"] = true
	}

	fmt.Println("IsLogin:",this.IsLogin,this.Loginuser)
	this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)
	this.TplName = "home.html"
}