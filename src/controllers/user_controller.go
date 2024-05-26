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

	usersRoute.Get("/", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.GetAllUsers)
	usersRoute.Get("/self", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.GetSelfUser)
	usersRoute.Get("/:id", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.GetUserById)
	usersRoute.Patch("/", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.UpdateUserImage)
	usersRoute.Post("/recover", services.RecoverPassword)
	usersRoute.Patch("/:id", auth.AuthMiddleware([]string{"Admin"}), services.UpdateUser)
	usersRoute.Delete("/:id", auth.AuthMiddleware([]string{"Admin"}), services.DeleteUser)
	// usersRoute.Patch("/", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.UpadteSelfUser)
	usersRoute.Delete("/", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.DeleteSelfUser)
}
