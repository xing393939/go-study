package app

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	//dsn := "root:123456@tcp(127.0.0.1:3306)?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "file:sqlite.db?cache=shared&mode=memory"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
