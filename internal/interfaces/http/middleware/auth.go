package middleware

import (
	"strings"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/common"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/infrastructure/config"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(common.NewErrorResponse(fiber.StatusUnauthorized, "Missing authorization header"))
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(common.NewErrorResponse(fiber.StatusUnauthorized, "Invalid authorization header format"))
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(config.AppConfig.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(common.NewErrorResponse(fiber.StatusUnauthorized, "Invalid or expired token"))
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(common.NewErrorResponse(fiber.StatusUnauthorized, "Invalid token claims"))
		}

		// Set user_id in context for handlers to use
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(common.NewErrorResponse(fiber.StatusUnauthorized, "Invalid user ID in token"))
		}

		c.Locals("user_id", uint(userIDFloat))

		return c.Next()
	}
}
