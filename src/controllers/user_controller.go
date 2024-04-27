package controllers

import (
	"pet_api/src/auth"
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func UserController(api fiber.Router) {
	usersRoute := api.Group("/users")

	usersRoute.Post("/register", services.CreateUser)
	usersRoute.Post("/login", services.LoginUser)

	usersRoute.Get("/", auth.AuthMiddleware([]string{"ONG", "Adoptador"}), services.GetAllUsers)
	usersRoute.Get("/:id", auth.AuthMiddleware([]string{"ONG", "Adoptador"}), services.GetUserById)
	usersRoute.Patch("/:id/img", auth.AuthMiddleware([]string{"ONG", "Adoptador"}), services.UpdateUserImage)
	usersRoute.Post("/recover_password", auth.AuthMiddleware([]string{"ONG", "Adoptador"}), services.RecoverPassword)
	usersRoute.Patch("/:id", auth.AuthMiddleware([]string{"ONG", "Adoptador"}), services.UpdateUser)
	usersRoute.Delete("/:id", auth.AuthMiddleware([]string{"ONG", "Adoptador"}), services.DeleteUser)
}
