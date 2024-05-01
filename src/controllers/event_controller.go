package controllers

import (
	"pet_api/src/auth"
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func EventController(api fiber.Router) {
	eventRoute := api.Group("/events")

	eventRoute.Get("/", auth.AuthMiddleware([]string{"ONG", "Admin"}), services.GetAllEvents)
	// eventRoute.Post("/", auth.AuthMiddleware([]string{"ONG", "Admin"}), services.CreateVisit)
	eventRoute.Get("/:id", auth.AuthMiddleware([]string{"ONG", "Admin"}), services.GetEventById)
	eventRoute.Delete("/:id", auth.AuthMiddleware([]string{"ONG", "Admin"}), services.DeleteEvent)
}
