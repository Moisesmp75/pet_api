package controllers

import (
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func UserController(api fiber.Router) {
	usersRoute := api.Group("/user")

	usersRoute.Get("/", services.GetAllUsers)
	usersRoute.Get("/:id", services.GetUserById)
	usersRoute.Post("/register", services.CreateUser)
	usersRoute.Post("/login", services.LoginUser)
	usersRoute.Patch("/:id", services.UpdateUserImage)
}
