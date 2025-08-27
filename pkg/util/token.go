package util

import (
	"time"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * time.Duration(config.AppConfig.JWT.ExpirationMinutes)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWT.Secret))
}
