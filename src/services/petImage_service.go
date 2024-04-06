package services

import (
	"log"
	"mime/multipart"
	"pet_api/src/common"
	"pet_api/src/models"
	"pet_api/src/repositories"
)

func CreatePetImages(petID uint, form *multipart.Form) ([]models.PetImage, error) {
	images, err := common.ImagesToBase64(form)
	if err != nil {
		return nil, err
	}

	var petImages []models.PetImage
	for _, base64Image := range images {
		petImage := models.PetImage{
			URL:   base64Image,
			PetID: petID,
		}

		createdImage, err := repositories.CreatePetImage(petImage)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		petImages = append(petImages, createdImage)
	}

	return petImages, nil
}