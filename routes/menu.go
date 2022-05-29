package routes

import (
	"github.com/galifornia/go-restaurant-management/models"
	"github.com/gofiber/fiber/v2"
)

func setupMenuRoutes(router fiber.Router) {
	menuApi := router.Group("/menu")
	menuApi.Get("/", getMenu)
}

func getMenu(c *fiber.Ctx) error {
	var menu models.Menu
	// TODO: call database
	return c.JSON(fiber.Map{"menu": menu})
}
