package main


import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/Go-SQL-Driver/MySQL"
)


var db *sql.DB

func main() {
	InitMysql()	
}


func InitMysql() {
	fmt.Println("InitMysql......")
	if db == nil {
		//  建立连接数据库的信息
		db, _ = sql.Open("mysql", "root:huchen12345677@tcp(127.0.0.1:3306)/myblogweb")
		CreateTableWithUser()
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