package routes

import (
	"github.com/galifornia/go-restaurant-management/controllers"
	"github.com/galifornia/go-restaurant-management/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupTableRoutes(router fiber.Router) {
	restricted := router.Group("/table")
	restricted.Use(middleware.SecureAuth())

	restricted.Get("/", controllers.GetAllTables)
	restricted.Get("/:id", controllers.GetTable)
	restricted.Post("/", controllers.NewTable)
	restricted.Patch("/:id", controllers.UpdateTable)
	restricted.Delete("/:id", controllers.DeleteTable)
}
