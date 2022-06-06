package models

import (
    "myblogweb/utils"
)

 type Article struct {
	Id int
	Title string
	Author string
	Tags string
	Short string
	Content string
	Createtime int64
 }

//  添加一篇文章
 func AddArticle(article Article) (int64,error) {
	i,err := insertArticle(article)
	return i,err
 }

//---------------数据库操作--------------------//

 //  插入一篇文章

func insertArticle(article Article) (int64,error) {
	return utils.ModifyDB("insert into article(title,tags,short,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title,article.Tags,article.Short,article.Content,article.Author,article.Createtime)
}


