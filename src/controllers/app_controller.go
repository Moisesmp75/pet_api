package controllers

import "github.com/gofiber/fiber/v2"

func AddControllers(api fiber.Router) {
	UserController(api)
	PetController(api)
	VisitController(api)
	ReportController(api)
}
