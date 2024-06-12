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
	donationRoute.Patch("/products/:id", auth.AuthMiddleware([]string{"ONG", "Admin"}), services.UpdateDonationProduct)

	donationRoute.Get("/money", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.GetAllDonationsMoney)
	donationRoute.Post("/money", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.CreateDonationMoney)
	donationRoute.Get("/money/:id", auth.AuthMiddleware([]string{"ONG", "Adoptador", "Duenio", "Admin"}), services.GetDonationMoneyById)
	donationRoute.Patch("/money/:id", auth.AuthMiddleware([]string{"ONG", "Admin"}), services.UpdateDonationMoney)
}
