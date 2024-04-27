package auth

import (
	"os"
	"pet_api/src/dto/response"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user response.LoginResponse) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role,
		"iat":   user.Iat.Unix(),
		"exp":   user.Exp.Unix(),
	})

	jwt_secret := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(jwt_secret))
}

func ParseJWTToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return claims, jwt.ErrInvalidKey
	}
	jwt_secret := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(parts[1], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwt_secret), nil
	})
	if err != nil {
		return claims, err
	}

	if !token.Valid {
		return claims, jwt.ErrSignatureInvalid
	}

	expClaims, ok := claims["exp"].(float64)
	if !ok {
		return claims, jwt.ErrInvalidKey
	}

	expirationTime := time.Unix(int64(expClaims), 0)
	if time.Now().After(expirationTime) {
		return claims, jwt.ErrTokenExpired
	}

	return claims, nil
}
