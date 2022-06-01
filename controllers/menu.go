package controllers

import (
	"github.com/galifornia/go-restaurant-management/database"
	"github.com/galifornia/go-restaurant-management/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllMenus(c *fiber.Ctx) error {
	var menus []models.Menu
	database.DB.Find(&menus).Limit(10)
	return c.JSON(menus)
}

func GetMenu(c *fiber.Ctx) error {
	var menu models.Menu
	id := c.Params("id")

	database.DB.First(&menu, "id = ?", id)
	return c.JSON(fiber.Map{"menu": menu})
}

func AddItemToMenu(c *fiber.Ctx) error {
	var menu models.Menu
	// TODO: call database
	return c.JSON(fiber.Map{"menu": menu})
}

func RemoveItemFromMenu(c *fiber.Ctx) error {
	var menu models.Menu
	// TODO: call database
	return c.JSON(fiber.Map{"menu": menu})
}
