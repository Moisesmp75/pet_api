package migration

import (
	"errors"
	"log"
	"pet_api/src/database"
	"pet_api/src/models"

	"gorm.io/gorm"
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
		var existingRole models.Role
		result := database.DB.Model(&models.Role{}).Where("name = ?", role.Name).First(&existingRole)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}

		if result.RowsAffected == 0 {
			if err := database.DB.Create(&role).Error; err != nil {
				return err
			}
		} else {
			if existingRole.Description != role.Description {
				existingRole.Description = role.Description
				if err := database.DB.Save(&existingRole).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}