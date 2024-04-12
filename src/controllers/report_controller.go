package controllers

import (
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func ReportController(api fiber.Router) {
	reportRoute := api.Group("/report")

	reportRoute.Get("/", services.GetAllReports)
	reportRoute.Post("/", services.CreateReport)
	reportRoute.Get("/:id", services.GetReportById)
}
