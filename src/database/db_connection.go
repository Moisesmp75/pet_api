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
	host := os.Getenv("DB_HOST")
	MYSQL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, db_name)
	var err error
	DB, err = gorm.Open(mysql.Open(MYSQL), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Cannot connect to Database")
	}
	log.Println("Connected to database.")
}
