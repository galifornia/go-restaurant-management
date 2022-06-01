package routes

import (
	"github.com/galifornia/go-restaurant-management/controllers"
	"github.com/galifornia/go-restaurant-management/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupNoteRoutes(router fiber.Router) {
	restricted := router.Group("/note")
	restricted.Use(middleware.SecureAuth())

	restricted.Get("/", controllers.GetAllNotes)
	restricted.Get("/:id", controllers.GetNote)
	restricted.Post("/", controllers.NewNote)
	restricted.Patch("/:id", controllers.UpdateNote)
	restricted.Delete("/:id", controllers.DeleteNote)
}
