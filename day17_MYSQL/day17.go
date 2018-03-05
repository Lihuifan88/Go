package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/testinfo?charset=utf8")
	if err != nil {
		panic(err)
	}
	//	stmt, err := db.Prepare("INSERT user_info SET userName=? ,departName=? , createTime=?")
	//	res, err := stmt.Exec("张三", "技术部", "2018-01-18")

	rows, err := db.Query("SELECT * FROM user_info")
	//	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	//	fmt.Print(id)

	for rows.Next() {
		var id int
		var userName string
		var departName string
		var createTime string
		rows.Scan(&id, &userName, &departName, &createTime)
		fmt.Print(id, userName, departName, createTime)
	}

}
