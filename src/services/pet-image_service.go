package services

import (
	"log"
	"mime/multipart"
	"pet_api/src/helpers"
	"pet_api/src/models"
	"pet_api/src/repositories"
)

func CreatePetImages(petID uint64, form *multipart.Form) ([]models.PetImage, error) {
	images, filenames, err := helpers.UploadFiles(form, "pet_images/")
	if err != nil {
		return nil, err
	}

	numFiles := len(images)
	if len(filenames) < numFiles {
		numFiles = len(filenames)
	}

	var petImages []models.PetImage
	for i := 0; i < numFiles; i++ {
		petImage := models.PetImage{
			Filename: filenames[i],
			URL:      images[i],
			PetID:    petID,
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
