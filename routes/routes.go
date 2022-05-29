package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	setupFoodRoutes(api)
	setupInvoiceRoutes(api)
	setupNoteRoutes(api)
	setupMenuRoutes(api)
	setupOrderRoutes(api)
	setupOrderItemRoutes(api)
	setupTableRoutes(api)
	setupUserRoutes(api)
}
