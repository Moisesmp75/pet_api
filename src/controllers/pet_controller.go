package controllers

import (
	"pet_api/src/auth"
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func PetController(api fiber.Router) {
	petsRoute := api.Group("/pets")

	petsRoute.Get("/", services.GetAllPets)
	petsRoute.Get("/:id", services.GetPetById)
	petsRoute.Post("/", auth.AuthMiddleware([]string{"ONG", "Duenio", "Admin"}), services.CreatePet)
	petsRoute.Patch("/:id", auth.AuthMiddleware([]string{"ONG", "Duenio", "Admin"}), services.UpdatePet)
	petsRoute.Delete("/:id", auth.AuthMiddleware([]string{"ONG", "Duenio", "Admin"}), services.DeletePet)
	petsRoute.Get("/types", auth.AuthMiddleware([]string{"ONG", "Duenio", "Admin", "Adoptador"}), services.GetAllPetTypes)
}
