package routes

import (
	"github.com/galifornia/go-restaurant-management/controllers"
	"github.com/galifornia/go-restaurant-management/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupFoodRoutes(router fiber.Router) {
	foodApi := router.Group("/food")

	foodApi.Get("/", controllers.GetAllFood)
	foodApi.Get("/:id", controllers.GetFoodDetails)

	restricted := router.Group("/food")
	restricted.Use(middleware.SecureAuth())

	restricted.Post("/", controllers.NewFood)
	restricted.Patch("/:id", controllers.UpdateFoodDetails)
	restricted.Delete("/:id", controllers.DeleteFood)
}
