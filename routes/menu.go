package routes

import (
	"github.com/galifornia/go-restaurant-management/controllers"
	"github.com/galifornia/go-restaurant-management/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupMenuRoutes(router fiber.Router) {
	menuApi := router.Group("/menu")

	menuApi.Get("/", controllers.GetMenu)

	restricted := router.Group("/menu")
	restricted.Use(middleware.SecureAuth())

	restricted.Post("/", controllers.AddItemToMenu)
	restricted.Delete("/:id", controllers.RemoveItemFromMenu)
}
