package controllers

import (
	"fmt"
	"myblogweb/models"
	"log"
)

type DeleteArticleController struct {
	BaseController
}


func (this *DeleteArticleController) Get() {
	artId,_ := this.GetInt("id")
	fmt.Println("删除id:",artId)
	_,err := models.DeleteArticle(artId)
	if err != nil {
		log.Println(err)
	}
	this.Redirect("/",302)
}