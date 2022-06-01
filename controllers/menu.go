package controllers

import (
	"github.com/galifornia/go-restaurant-management/models"
	"github.com/gofiber/fiber/v2"
)

func GetMenu(c *fiber.Ctx) error {
	var menu models.Menu
	// TODO: call database
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
