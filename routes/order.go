package routes

import (
	"github.com/galifornia/go-restaurant-management/controllers"
	"github.com/galifornia/go-restaurant-management/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupOrderRoutes(router fiber.Router) {
	restricted := router.Group("/order")
	restricted.Use(middleware.SecureAuth())

	restricted.Get("/", controllers.GetAllOrders)
	restricted.Get("/:id", controllers.GetOrder)
	restricted.Post("/", controllers.NewOrder)
	restricted.Patch("/:id", controllers.UpdateOrder)
	restricted.Delete("/:id", controllers.DeleteOrder)
}
