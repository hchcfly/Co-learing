package main

import (
	"database/sql"
	"fmt"
)


func InitMysql() {
	fmt.Println("InitMysql....")
	if db == nil {
		sql.Open("mysql","root:hchc32@tcp(127.0.0.1:3306)/myblogweb")
		
	}

}


func main() {
	
}