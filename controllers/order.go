package controllers

import (
	"github.com/galifornia/go-restaurant-management/database"
	"github.com/galifornia/go-restaurant-management/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllOrders(c *fiber.Ctx) error {
	var orders []models.Order
	database.DB.Find(&orders)
	return c.JSON(orders)
}

func GetOrder(c *fiber.Ctx) error {
	var order models.Order
	id := c.Params("id")

	database.DB.First(&order, "uuid = ?", id)
	if order.UUID == "" {
		return c.Status(500).SendString("Could not find order with provided 'id'")
	}

	return c.JSON(order)
}

func NewOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		c.Status(500).SendString("Could not parse body data from POST request")
	}

	order.UUID = uuid.NewString()
	database.DB.Create(&order)
	return c.JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	var order models.Order
	database.DB.First(&order, "uuid = ?", id)
	if order.UUID == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error"})
	}

	database.DB.Delete(&order)
	return c.Status(200).JSON(fiber.Map{"status": "successful"})
}

func UpdateOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	var order models.Order
	database.DB.First(&order, "uuid = ?", id)
	if order.UUID == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error"})
	}

	var updatedOrder models.Order

	err := c.BodyParser(&updatedOrder)
	if err != nil {
		return c.Status(503).JSON(err.Error())
	}

	database.DB.Model(&order).Updates(updatedOrder)

	return c.Status(200).JSON(updatedOrder)
}
