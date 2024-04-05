package auth

import (
	"os"
	"pet_api/src/dto/response"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user response.LoginResponse) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    user.Email,
		"role":     user.Role,
		"iat":      user.Iat.Unix(),
		"exp":      user.Exp.Unix(),
	})

	jwt_secret := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(jwt_secret))
}
