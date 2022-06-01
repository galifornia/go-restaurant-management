package controllers

import (
	"github.com/galifornia/go-restaurant-management/database"
	"github.com/galifornia/go-restaurant-management/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewMenu(c *fiber.Ctx) error {
	var menu models.Menu

	if err := c.BodyParser(&menu); err != nil {
		c.Status(500).SendString("Could not parse body data from POST request")
	}

	menu.UUID = uuid.NewString()
	database.DB.Create(&menu)
	return c.JSON(menu)
}

func DeleteMenu(c *fiber.Ctx) error {
	id := c.Params("id")

	var menu models.Menu
	database.DB.First(&menu, "uuid = ?", id)
	if menu.Name == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error"})
	}

	database.DB.Delete(&menu)
	return c.Status(200).JSON(fiber.Map{"status": "successful"})
}

func GetAllMenus(c *fiber.Ctx) error {
	var menus []models.Menu
	database.DB.Find(&menus).Limit(10)
	return c.JSON(menus)
}

func GetMenu(c *fiber.Ctx) error {
	var menu models.Menu
	id := c.Params("id")

	database.DB.First(&menu, "uuid = ?", id)
	if menu.Name == "" {
		return c.Status(500).SendString("Could not find food with provided 'id'")
	}
	return c.JSON(menu)

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
