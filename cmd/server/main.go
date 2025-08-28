package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/application"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/infrastructure/config"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/infrastructure/database"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/infrastructure/logging"
	infra_repo "github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/infrastructure/repository"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/interfaces/http/handler"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/interfaces/http/middleware"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/interfaces/http/router"
	"github.com/gofiber/fiber/v3"
)

func main() {
	config.LoadConfig()
	cfg := &config.AppConfig
	logger := logging.NewAsyncLogger(cfg.Log.FilePath)
	logger.Info("Configuration and logger initialized")

	db, err := database.NewMySQLConnection(cfg)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	logger.Info("Database connection established")

	// Repositories
	userRepository := infra_repo.NewGormUserRepository(db)
	contactRepository := infra_repo.NewGormContactRepository(db)
	addressRepository := infra_repo.NewGormAddressRepository(db)

	// Services
	authService := application.NewAuthService(userRepository, logger)
	userService := application.NewUserService(userRepository, logger)
	contactService := application.NewContactService(contactRepository, userRepository, logger)
	addressService := application.NewAddressService(addressRepository, contactRepository, logger)

	// Handlers
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	contactHandler := handler.NewContactHandler(contactService)
	addressHandler := handler.NewAddressHandler(addressService)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			// Default status code adalah 500
			code := fiber.StatusInternalServerError
			message := "Terjadi kesalahan internal pada server"

			// Cek apakah error-nya adalah tipe fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				message = e.Message
			}

			// Kirim response JSON yang konsisten
			return c.Status(code).JSON(fiber.Map{
				"code":   code,
				"status": "Error",
				"errors": message,
			})
		},
	})

	app.Use(middleware.Logger(logger))

	router.SetupRoutes(app, authHandler, userHandler, contactHandler, addressHandler)

	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	logger.Infof("Starting server on %s", serverAddr)
	if err := app.Listen(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
