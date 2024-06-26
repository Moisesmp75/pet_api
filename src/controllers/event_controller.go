package controllers

import (
	"pet_api/src/auth"
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func EventController(api fiber.Router) {
	eventRoute := api.Group("/events")

	eventRoute.Get("/", auth.AuthMiddleware([]string{"ONG", "Duenio", "Admin", "Adoptador"}), services.GetAllEvents)
	eventRoute.Post("/", auth.AuthMiddleware([]string{"ONG", "Admin"}), services.CreateEvent)
	eventRoute.Get("/:id", auth.AuthMiddleware([]string{"ONG", "Duenio", "Admin", "Adoptador"}), services.GetEventById)
	eventRoute.Delete("/:id", auth.AuthMiddleware([]string{"ONG", "Admin"}), services.DeleteEvent)
}
