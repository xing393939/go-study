package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func ConnectDatabase(dsn string, maxIdleConns int, maxOpenConns int) (*gorm.DB, *sql.DB) {
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务
		PrepareStmt:            true, // 创建并缓存预编译语句
	})

	if err != nil {
		fmt.Printf("gormDB.Setup err: %v", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		fmt.Printf("gormDB.Setup err: %v", err)
	}
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Minute * 60)
	return gormDB, sqlDB
}

func testExecTimeout2() {
	dsn := "root:123456@tcp(localhost:3306)/test?timeout=1s"
	db, sqlDb := ConnectDatabase(dsn, 1, 2)

	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		stats := sqlDb.Stats()
		str := fmt.Sprintf("%+v", stats)
		w.Write([]byte(str))
	})
	http.HandleFunc("/db", func(w http.ResponseWriter, r *http.Request) {
		var a int64
		db.Debug().Raw("SELECT version()").Scan(&a)
		w.Write([]byte("db"))
	})
	go func() {
		for {
			db.Debug().Exec("SELECT sleep(60)")
		}
	}()
	http.ListenAndServe("0.0.0.0:6060", nil)
}

func testConnTimeout(addr string, timeout int, suffix string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	dsn := fmt.Sprintf("root:123456@%s/test?timeout=%ds%s", addr, timeout, suffix)
	ConnectDatabase(dsn, 1, 2)
}

func testExecTimeout() {
	dsn := "root:123456@tcp(43.192.68.64:3306)/test?readTimeout=5s"
	db, sqlDb := ConnectDatabase(dsn, 1, 2)
	db.Debug().Exec("SELECT sleep(5)")
	sqlDb.Stats()
}

func main() {
	// 一，i/o timeout错误：iptables -I OUTPUT -p tcp --sport 3306 --tcp-flags SYN SYN -j DROP
	// 3306端口没有响应
	// testConnTimeout("tcp(43.192.68.64:3306)", 3, "")

	// 二，bad connection错误：stop mysql && ncat --broker --listen -p 3306
	// 3306端口虽然可以连接，但是不是mysql协议
	// testConnTimeout("tcp(43.192.68.64:3306)", 3, "&readTimeout=5s")

	// 三，执行sql的超时时间，readTimeout和writeTimeout
	// testExecTimeout()
}
