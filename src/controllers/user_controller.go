package controllers

import (
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func UserController(api fiber.Router) {
	usersRoute := api.Group("/user")

	usersRoute.Get("/", services.GetAllUsers)
	usersRoute.Post("/", services.CreateUser)
}
