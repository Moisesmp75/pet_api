package controllers

import (
	"pet_api/src/auth"
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func ReportController(api fiber.Router) {
	reportRoute := api.Group("/reports")

	reportRoute.Get("/", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.GetAllReports)
	reportRoute.Post("/", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.CreateReport)
	reportRoute.Get("/:id", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.GetReportById)
}
