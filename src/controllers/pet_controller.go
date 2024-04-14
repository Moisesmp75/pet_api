package controllers

import (
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func PetController(api fiber.Router) {
	petsRoute := api.Group("/pets")

	petsRoute.Get("/", services.GetAllPets)
	petsRoute.Get("/:id", services.GetPetById)
	petsRoute.Post("/", services.CreatePet)
	petsRoute.Patch("/:id/img", services.UpdatePetImages)
	petsRoute.Patch("/:id", services.UpdatePet)
}
