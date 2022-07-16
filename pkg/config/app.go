package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "ahmed:root@tcp(127.0.0.1:3306)/bookstore-crud?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Println("Connection Failed to Open")
		panic(err)
	}
	db = d
	log.Println("Connection Established")
}

func GetDB() *gorm.DB {
	return db
}
