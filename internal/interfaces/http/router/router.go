package router

import (
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/interfaces/http/handler"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/interfaces/http/middleware"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(
	app *fiber.App,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	contactHandler *handler.ContactHandler,
	addressHandler *handler.AddressHandler,
) {
	api := app.Group("/api")

	// Auth routes (Public)
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	// User routes (Protected)
	users := api.Group("/users", middleware.Auth())
	users.Get("/current", userHandler.GetCurrent)
	users.Get("/", userHandler.GetAll)
	users.Get("/:id", userHandler.GetByID)
	users.Put("/:id", userHandler.Update)
	users.Delete("/:id", userHandler.Delete)
	// Contact related to user
	users.Post("/:userId/contacts", contactHandler.Create)
	users.Get("/:userId/contacts", contactHandler.GetAllByUserID)

	// Contact routes (Protected)
	contacts := api.Group("/contacts", middleware.Auth())
	contacts.Get("/:id", contactHandler.GetByID)
	contacts.Put("/:id", contactHandler.Update)
	contacts.Delete("/:id", contactHandler.Delete)
	// Address related to contact
	contacts.Post("/:contactId/addresses", addressHandler.Create)
	contacts.Get("/:contactId/addresses", addressHandler.GetAllByContactID)

	// Address routes (Protected)
	addresses := api.Group("/addresses", middleware.Auth())
	addresses.Get("/:id", addressHandler.GetByID)
	addresses.Put("/:id", addressHandler.Update)
	addresses.Delete("/:id", addressHandler.Delete)
}
