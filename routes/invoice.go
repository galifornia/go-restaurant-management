package routes

import (
	"github.com/galifornia/go-restaurant-management/controllers"
	"github.com/galifornia/go-restaurant-management/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupInvoiceRoutes(router fiber.Router) {
	invoiceApi := router.Group("/invoice")

	invoiceApi.Get("/", controllers.GetAllInvoices)
	invoiceApi.Get("/:id", controllers.GetInvoice)

	restricted := router.Group("/invoice")
	restricted.Use(middleware.SecureAuth())

	restricted.Post("/", controllers.NewInvoice)
	restricted.Patch("/:id", controllers.UpdateInvoiceDetails)
	restricted.Delete("/:id", controllers.DeleteInvoice)
}
