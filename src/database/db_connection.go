package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDatabase() {
	var err error
	const MYSQL = "root:1234@tcp(127.0.0.1:3306)/pet_db?charset=utf8mb4&parseTime=True&loc=Local"
	const DSN = MYSQL
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Cannot connect to Database")
	}
	log.Println("Connected to database.")
}
