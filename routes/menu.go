package routes

import (
	"github.com/galifornia/go-restaurant-management/controllers"
	"github.com/galifornia/go-restaurant-management/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupMenuRoutes(router fiber.Router) {
	menuApi := router.Group("/menu")

	menuApi.Get("/", controllers.GetAllMenus)
	menuApi.Get("/:id", controllers.GetMenu)

	restricted := router.Group("/menu")
	restricted.Use(middleware.SecureAuth())

	restricted.Post("/", controllers.NewMenu)
	restricted.Delete("/", controllers.DeleteMenu)
	restricted.Post("/:id", controllers.AddItemToMenu)
	restricted.Delete("/:id", controllers.RemoveItemFromMenu)
}
