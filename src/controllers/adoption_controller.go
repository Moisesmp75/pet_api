package controllers

import (
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func AdoptionController(api fiber.Router) {
	adoptionRoute := api.Group("/adoption")

	adoptionRoute.Get("/", services.GetAllAdoptions)
	adoptionRoute.Post("/", services.CreateAdoption)
	adoptionRoute.Get("/:id", services.GetAdoptionById)
}
