package utils

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"

)

var db *sql.DB

func InitMysql() {
	fmt.Println("InitMysql......")
	if db == nil {
		//  建立连接数据库的信息
		db, _ = sql.Open("mysql", "root:huchen12345677@tcp(127.0.0.1:3306)/myblogweb")
		CreateTableWithUser()
		CreateTableWithArticle()
		fmt.Printf("db: %v\n", db)
	}
}
//  操作数据库
func ModifyDB(sql string,args ...interface{}) (int64,error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0,err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0,err
	}
	return count,nil
}

//  创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
	);`
	ModifyDB(sql)
}

//  创建文章表
func CreateTableWithArticle() {
	sql := `CREATE TABLE IF NOT EXISTS article(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title VARCHAR(30),
		author VARCHAR(20),
		tags VARCHAR(30),
		short VARCHAR(255),
		content LONGTEXT,
		createtime INT(10)
		);`
	ModifyDB(sql)

}

//  查询文章表
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}