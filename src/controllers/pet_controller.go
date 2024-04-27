package controllers

import (
	"pet_api/src/auth"
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func PetController(api fiber.Router) {
	petsRoute := api.Group("/pets")

	petsRoute.Use(auth.AuthMiddleware([]string{"ONG", "Adoptador"}))

	petsRoute.Get("/", auth.AuthMiddleware([]string{"ONG", "Adoptador"}), services.GetAllPets)
	petsRoute.Get("/:id", auth.AuthMiddleware([]string{"ONG", "Adoptador"}), services.GetPetById)
	petsRoute.Post("/", auth.AuthMiddleware([]string{"ONG"}), services.CreatePet)
	petsRoute.Patch("/:id/img", auth.AuthMiddleware([]string{"ONG"}), services.UpdatePetImages)
	petsRoute.Patch("/:id", auth.AuthMiddleware([]string{"ONG"}), services.UpdatePet)
	petsRoute.Delete("/:id", auth.AuthMiddleware([]string{"ONG"}), services.DeletePet)
}
