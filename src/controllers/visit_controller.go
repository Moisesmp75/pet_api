package controllers

import (
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func VisitController(api fiber.Router) {
	visitRoute := api.Group("/visit")

	visitRoute.Get("/", services.GetAllVisits)
	visitRoute.Post("/", services.CreateVisit)
	visitRoute.Get("/:id", services.GetVisitById)
}