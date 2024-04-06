package controllers

import (
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func PetController(api fiber.Router) {
	petsRoute := api.Group("/pet")

	petsRoute.Get("/", services.GetAllPets)
	petsRoute.Get("/:id", services.GetPetById)
	petsRoute.Post("/", services.CreatePet)
	petsRoute.Post("/:id/img", services.UpdatePetImages)
}
