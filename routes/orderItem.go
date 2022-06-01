package routes

import (
	"github.com/galifornia/go-restaurant-management/controllers"
	"github.com/galifornia/go-restaurant-management/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupOrderItemRoutes(router fiber.Router) {
	restricted := router.Group("/order-item")
	restricted.Use(middleware.SecureAuth())

	restricted.Get("/:id", controllers.GetOrderItem)
	restricted.Post("/", controllers.NewOrderItem)
	restricted.Patch("/:id", controllers.UpdateOrderItem)
	restricted.Delete("/:id", controllers.DeleteOrderItem)
}
