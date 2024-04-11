package services

import (
	"log"
	"mime/multipart"
	"pet_api/src/helpers"
	"pet_api/src/models"
	"pet_api/src/repositories"
)

func CreatePetImages(petID uint64, form *multipart.Form) ([]models.PetImage, error) {
	images, err := helpers.UploadFiles(form, "pet_images/")
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
