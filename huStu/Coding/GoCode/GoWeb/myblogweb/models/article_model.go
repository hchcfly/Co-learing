package models

import (
	"myblogweb/utils"
	"fmt"
	"github.com/astaxie/beego"
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

 //  查询文章
 //  根据页码查询文章
 func FindArticleWithPage(page int) ([]Article ,error) {
	// 从配置文件中获取每一页的文章数量
	num,_ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("------------------>page",page)
	return QueryArticleWithPage(page,num)
 }

//---------------数据库操作--------------------//

 //  插入一篇文章

func insertArticle(article Article) (int64,error) {
	return utils.ModifyDB("insert into article(title,tags,short,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title,article.Tags,article.Short,article.Content,article.Author,article.Createtime)
}

// 分页查询数据库1
func QueryArticleWithPage(page ,num int) ([]Article,error) {
	sql := fmt.Sprintf("limit %d,%d",page*num,num)
	return QueryArticleWithCon(sql)
}
// 分页查询数据库2
func QueryArticleWithCon(sql string) ([]Article,error) {
	sql = "select id,title,tags,short,content,author,createtime from article " + sql
	fmt.Println("sql:",sql)
	rows,err := utils.QueryDB(sql)
	if err != nil {
		return nil,err
	}
	var artList []Article
	for rows.Next() {
		id := 0
        title := ""
        tags := ""
        short := ""
        content := ""
		author := ""
		var createtime int64
        createtime = 0
        rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
        art := Article{id, title, tags, short, content, author, createtime}
        artList = append(artList, art)
	}
	return artList,nil
}




