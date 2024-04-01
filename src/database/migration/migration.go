package migration

import (
	"log"
	"pet_api/src/database"
	"pet_api/src/models"
)

func RunMigration() {
	errUser := database.DB.AutoMigrate(&models.User{})
	if errUser != nil {
		log.Fatal(errUser.Error())
	}
}
