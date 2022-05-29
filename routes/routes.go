package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// UNRESTRICTED ROUTES
	setupFoodRoutes(api)
	setupInvoiceRoutes(api)
	setupNoteRoutes(api)
	setupUserUnrestrictedRoutes(api)
	setupMenuRoutes(api)

	// JWT Middleware
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	// RESTRICTED ROUTES
	setupUserRoutes(api)
	setupOrderRoutes(api)
	setupOrderItemRoutes(api)
	setupTableRoutes(api)
}
