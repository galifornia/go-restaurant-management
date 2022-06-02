package routes

import (
	"github.com/galifornia/go-restaurant-management/controllers"
	"github.com/galifornia/go-restaurant-management/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupFoodMenuRoutes(router fiber.Router) {
	restricted := router.Group("/food-menu")
	restricted.Use(middleware.SecureAuth())

	restricted.Get("/:id", controllers.GetFoodsInMenu)
	restricted.Post("/", controllers.AddFoodToMenu)
	restricted.Delete("/:id", controllers.RemoveFoodFromMenu)
}
