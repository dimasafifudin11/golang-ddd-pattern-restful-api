package router

import (
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/interfaces/http/handler"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/interfaces/http/middleware"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, authHandler *handler.AuthHandler, userHandler *handler.UserHandler) {
	api := app.Group("/api")

	// Auth routes (Public)
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	// User routes (Protected)
	users := api.Group("/users", middleware.Auth()) // Terapkan middleware di sini
	users.Get("/current", userHandler.GetCurrent)
}
