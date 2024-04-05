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

	errPet := database.DB.AutoMigrate(&models.Pet{})
	if errPet != nil {
		log.Fatal(errPet.Error())
	}

	errRol := database.DB.AutoMigrate(&models.Role{})
	if errRol != nil {
		log.Fatal(errRol.Error())
	}

	if err := SetupDefaultRoles(); err != nil {
		log.Fatal(err.Error())
	}
}

func SetupDefaultRoles() error {
	roles := []models.Role{
		{Name: "ONG", Description: "Organización No Gubernamental"},
		{Name: "Adoptador", Description: "Usuario que puede adoptar, dar en adopción y hacer donaciones"},
	}

	for _, role := range roles {
		if err := database.DB.Create(&role).Error; err != nil {
			return err
		}
	}

	return nil
}