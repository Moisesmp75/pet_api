package controllers

import (
	"pet_api/src/auth"
	"pet_api/src/services"

	"github.com/gofiber/fiber/v2"
)

func DonationController(api fiber.Router) {
	donationRoute := api.Group("/donations")

	donationRoute.Get("/products", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.GetAllDonationsProduct)
	donationRoute.Post("/products", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.CreateDonationProduct)
	donationRoute.Get("/products/:id", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.GetDonationProductById)
}
