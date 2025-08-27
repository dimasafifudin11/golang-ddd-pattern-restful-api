package main

import (
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
	// =========================================================================
	// Inisialisasi Awal (Fondasi Aplikasi)
	// =========================================================================

	// 1. Muat Konfigurasi dari file .yaml menggunakan Viper
	config.LoadConfig()
	cfg := &config.AppConfig
	logger := logging.NewAsyncLogger(cfg.Log.FilePath)
	logger.Info("Configuration and logger initialized")

	// 2. Buat Koneksi Database menggunakan GORM
	db, err := database.NewMySQLConnection(cfg)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	logger.Info("Database connection established")

	// =========================================================================
	// Dependency Injection (Merakit Komponen Aplikasi)
	// =========================================================================

	// 3. Inisialisasi Repository (Layer Akses Data)
	// Repository butuh koneksi database (db) untuk bisa bekerja.
	userRepository := infra_repo.NewGormUserRepository(db)

	// 4. Inisialisasi Service (Layer Logika Bisnis)
	// Service butuh Repository untuk mengakses data dan logger untuk mencatat log.
	authService := application.NewAuthService(userRepository, logger)
	userService := application.NewUserService(userRepository, logger)

	// 5. Inisialisasi Handler (Layer HTTP/Interface)
	// Handler butuh Service untuk menjalankan logika bisnis.
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)

	// =========================================================================
	// Setup & Menjalankan Web Server
	// =========================================================================

	// 6. Inisialisasi Aplikasi Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			// Custom error handler agar output error konsisten
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"code":   code,
				"status": "Error",
				"data":   err.Error(),
			})
		},
	})

	// 7. Daftarkan Middleware
	// Middleware Logger akan dijalankan untuk setiap request yang masuk.
	app.Use(middleware.Logger(logger))

	// 8. Daftarkan Semua Rute (Endpoint)
	// Kita serahkan pengaturan rute ke fungsi SetupRoutes, sambil memberikan handler yang sudah 'dirakit'.
	router.SetupRoutes(app, authHandler, userHandler)

	// 9. Jalankan Server
	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	logger.Infof("Starting server on %s", serverAddr)
	if err := app.Listen(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
