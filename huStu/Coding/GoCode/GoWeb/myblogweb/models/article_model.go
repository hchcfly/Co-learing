package models

import (
	"myblogweb/utils"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"log"
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
	SetArticleRowsNum()
	return i,err
 }

 //  查询文章
 //  根据页码查询文章
 func FindArticleWithPage(page int) ([]Article ,error) {
	// 从配置文件中获取每一页的文章数量
	num,_ := beego.AppConfig.Int("articleListPageNum")
	page--
	return QueryArticleWithPage(page,num)
 }

 //  根据文章id查询文章
 func QueryArticleWithId(id int) (Article ,error) {
	row := utils.QueryRowDB("select id,title,author,tags,short,content,createtime from article where id=" + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &author, &tags, &short, &content, &createtime)
	art := Article{id, title, author, tags, short, content, createtime}
	return art,nil
}


//更新文章
func UpdateArticle(article Article) (int64,error) {
	return utils.ModifyDB("update article set title=?,tags=?,short=?,content=? where id=?",
		article.Title, article.Tags, article.Short, article.Content, article.Id)
}


//删除文章
func DeleteArticle(artId int) (int64,error) {
	i,err := deleteArticleWithArtId(artId)
	SetArticleRowsNum()
	return i,err
}

//查询标签,返回一个字段的列表
func QueryArticleWithParam(param string) []string {
	//查询特定字段的一列值
	rows ,err := utils.QueryDB(fmt.Sprintf("select %s from article",param))
	if err!=nil {
		log.Println(err)
	}
	var paramList []string
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList,arg)
	}
	return paramList
}

//根据标签tag查询文章
func QueryArticleWithTag(tag string) ([]Article ,error) {
	sql := fmt.Sprintf("where tags = \"%s\"",tag)
	return QueryArticleWithCon(sql)
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
	sql = "select id,title,author,tags,short,content,createtime from article " + sql
	// fmt.Println("sql:",sql)
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
        rows.Scan(&id, &title, &author, &tags, &short, &content, &createtime)
		art := Article{id, title, author, tags, short, content, createtime}
        artList = append(artList, art)
	}
	return artList,nil
}

var articleRowNum = 0

//获取当前文章总数
func GetArticleRowNum() int {
	if articleRowNum == 0 {
		articleRowNum = QueryArticleRowNum()
	}
	return articleRowNum
}

//数据库查询当前文章总数
func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}

//从数据库中查询文章总数[add/delete]
func SetArticleRowsNum() {
	articleRowNum = QueryArticleRowNum()
}

//删除对应文章id的文章
func deleteArticleWithArtId(artId int) (int64,error) {
	return utils.ModifyDB("delete from article where id=?",artId)
}


