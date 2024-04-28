package controllers

import (
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func RoleController(api fiber.Router) {
	roleRoute := api.Group("/roles")

	roleRoute.Get("/", services.GetAllRoles)
}
