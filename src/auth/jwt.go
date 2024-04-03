package auth

import (
	"os"
	"pet_api/src/dto/response"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user response.UserResponse) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    user.Email,
		"username": user.Username,
	})

	jwt_secret := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(jwt_secret))
}
