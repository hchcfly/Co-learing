package controllers

import(
	"fmt"
	"myblogweb/models"
)

type TagsController struct {
	BaseController
}


func (this *TagsController)Get() {
	//查询标签列表
	tags := models.QueryArticleWithParam("tags")
	fmt.Println(models.HandleTagsListData(tags))
	this.Data["Tags"] = models.HandleTagsListData(tags)
	this.TplName = "tags.html"
}
