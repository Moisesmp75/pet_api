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

	errPetImg := database.DB.AutoMigrate(&models.PetImage{})
	if errPetImg != nil {
		log.Fatal(errPetImg.Error())
	}

	errPetType := database.DB.AutoMigrate(&models.PetType{})
	if errPetType != nil {
		log.Fatal(errPetType.Error())
	}

	if err := SetupDefaultPetTypes(); err != nil {
		log.Fatal(err.Error())
	}

	errVisit := database.DB.AutoMigrate(&models.Visit{})
	if errVisit != nil {
		log.Fatal(errVisit.Error())
	}

	errReport := database.DB.AutoMigrate(&models.Report{})
	if errReport != nil {
		log.Fatal(errReport.Error())
	}

	errAdoption := database.DB.AutoMigrate(&models.Adoption{})
	if errAdoption != nil {
		log.Fatal(errAdoption.Error())
	}

	errEvents := database.DB.AutoMigrate(&models.Event{})
	if errEvents != nil {
		log.Fatal(errEvents.Error())
	}

	errBankAccount := database.DB.AutoMigrate(&models.BankAccount{})
	if errBankAccount != nil {
		log.Fatal(errBankAccount.Error())
	}

	errONGInfo := database.DB.AutoMigrate(&models.ONGInfo{})
	if errONGInfo != nil {
		log.Fatal(errONGInfo.Error())
	}

	errUserImage := database.DB.AutoMigrate(&models.UserImage{})
	if errUserImage != nil {
		log.Fatal(errUserImage.Error())
	}

	errPetBehavior := database.DB.AutoMigrate(&models.PetBehavior{})
	if errPetBehavior != nil {
		log.Fatal(errPetBehavior.Error())
	}

	errProduct := database.DB.AutoMigrate(&models.Product{})
	if errProduct != nil {
		log.Fatal(errProduct.Error())
	}

	errDonationProd := database.DB.AutoMigrate(&models.DonationProduct{})
	if errDonationProd != nil {
		log.Fatal(errDonationProd.Error())
	}

	errDonationMoney := database.DB.AutoMigrate(&models.DonationMoney{})
	if errDonationMoney != nil {
		log.Fatal(errDonationMoney.Error())
	}
}

func SetupDefaultRoles() error {
	roles := []models.Role{
		{Name: "ONG", Description: "Organización No Gubernamental"},
		{Name: "Adoptador", Description: "Usuario que puede adoptar y hacer donaciones"},
		{Name: "Duenio", Description: "Usuario que puede dar en adopcion y hacer donaciones"},
		{Name: "Admin", Description: "Super usuario con todos los permisos"},
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

func SetupDefaultPetTypes() error {
	petTypes := []models.PetType{
		{Name: "Dog"},
		{Name: "Hamster"},
		{Name: "Cat"},
		{Name: "Rabbit"},
	}

	for _, petType := range petTypes {
		var existingType models.PetType
		result := database.DB.Where("name = ?", petType.Name).First(&existingType)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}

		if result.RowsAffected == 0 {
			if err := database.DB.Create(&petType).Error; err != nil {
				return err
			}
			log.Printf("Created default pet type: %s\n", petType.Name)
		}
	}

	return nil
}
