package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDatabase() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	MYSQL := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, port, db_name)
	var err error
	DB, err = gorm.Open(mysql.Open(MYSQL), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Cannot connect to Database")
	}
	log.Println("Connected to database.")
}
