// simple.go

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

func main() {
	a := http.Client{
		Timeout: time.Millisecond * 10,
	}
	res, err := a.Get("http://192.168.2.119:8008/answerfun/activityInfo?actuniqid=769e5c883b56")
	if err != nil {
		println(err.Error())
		return
	}
	println(res.StatusCode)
}
