package controllers

import (
	"pet_api/src/auth"
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func VisitController(api fiber.Router) {
	visitRoute := api.Group("/visits")

	visitRoute.Get("/", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.GetAllVisits)
	visitRoute.Post("/", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.CreateVisit)
	visitRoute.Get("/:id", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.GetVisitById)
	visitRoute.Delete("/:id", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.DeleteVisit)
}
