package auth

import (
	"fmt"
	"pet_api/src/dto/response"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(allowedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("Authorization Token not provided"))
		}

		claims, err := ParseJWTToken(authHeader)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("Invalid Token"))
		}
		fmt.Println(claims)
		userEmail, ok := claims["email"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("Error getting user_id from token"))
		}

		role, ok := claims["role"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("Error getting role from token"))
		}

		isAllowed := false
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("Access denied"))
		}
		c.Locals("user_email", userEmail)

		return c.Next()
	}
}
