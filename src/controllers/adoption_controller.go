package controllers

import (
	"pet_api/src/auth"
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func AdoptionController(api fiber.Router) {
	adoptionRoute := api.Group("/adoptions")

	adoptionRoute.Get("/", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.GetAllAdoptions)
	adoptionRoute.Post("/", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.CreateAdoption)
	adoptionRoute.Get("/:id", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.GetAdoptionById)
	adoptionRoute.Delete("/:id", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.DeleteAdoption)
	adoptionRoute.Patch("/:id", auth.AuthMiddleware([]string{"Duenio", "Admin", "ONG"}), services.UpdateAdoption)
}
