package controllers

import (
	"github.com/galifornia/go-restaurant-management/database"
	"github.com/galifornia/go-restaurant-management/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetTables(c *fiber.Ctx) error {
	var tables []models.Table
	database.DB.Find(&tables)
	return c.JSON(tables)
}

func GetTable(c *fiber.Ctx) error {
	var table models.Table
	id := c.Params("id")

	database.DB.First(&table, "uuid = ?", id)
	if table.UUID == "" {
		return c.Status(500).SendString("Could not find Table with provided 'id'")
	}

	return c.JSON(table)
}

func NewTable(c *fiber.Ctx) error {
	var table models.Table

	if err := c.BodyParser(&table); err != nil {
		c.Status(500).SendString("Could not parse body data from POST request")
	}

	table.UUID = uuid.NewString()
	database.DB.Create(&table)
	return c.JSON(table)
}

func DeleteTable(c *fiber.Ctx) error {
	id := c.Params("id")

	var table models.Table
	database.DB.First(&table, "uuid = ?", id)
	if table.UUID == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error"})
	}

	database.DB.Delete(&table)
	return c.Status(200).JSON(fiber.Map{"status": "successful"})
}

func UpdateTable(c *fiber.Ctx) error {
	id := c.Params("id")

	var table models.Table
	database.DB.First(&table, "uuid = ?", id)
	if table.UUID == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error"})
	}

	var updatedTable models.Table

	err := c.BodyParser(&updatedTable)
	if err != nil {
		return c.Status(503).JSON(err.Error())
	}

	database.DB.Model(&table).Updates(updatedTable)

	return c.Status(200).JSON(updatedTable)
}
