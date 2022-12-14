// simple.go

package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var dataBase = "root:Aa123456@tcp(192.168.0.101:3306)/?timeout=5s&readTimeout=6s"

func mysqlInit() {
	var err error
	DB, err = sql.Open("mysql", dataBase)
	if err != nil {
		log.Fatalln("open db fail:", err)
	}

	DB.SetMaxOpenConns(3)
	DB.SetMaxIdleConns(3)
}

func main() {
	mysqlInit()

	for {
		log.Println("start")
		execSql()
		time.Sleep(3 * time.Second)
	}
}

func execSql() {
	var value int
	err := DB.QueryRow("select 1").Scan(&value)
	if err != nil {
		log.Println("query failed:", err)
		return
	}

	log.Println("value:", value)
}
