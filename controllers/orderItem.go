package controllers

import (
	"github.com/galifornia/go-restaurant-management/database"
	"github.com/galifornia/go-restaurant-management/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllOrderItems(c *fiber.Ctx) error {
	var ordersItems []models.OrderItem
	database.DB.Find(&ordersItems)
	return c.JSON(ordersItems)
}

func GetOrderItem(c *fiber.Ctx) error {
	var orderItem models.OrderItem
	id := c.Params("id")

	database.DB.First(&orderItem, "uuid = ?", id)
	if orderItem.UUID == "" {
		return c.Status(500).SendString("Could not find OrderItem with provided 'id'")
	}

	return c.JSON(orderItem)
}

func GetOrderItemsByOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")

	var ordersItems []models.OrderItem
	database.DB.Find(&ordersItems, "order_id = ?", orderID)

	return c.JSON(ordersItems)
}

func NewOrderItem(c *fiber.Ctx) error {
	var orderItem models.OrderItem

	if err := c.BodyParser(&orderItem); err != nil {
		c.Status(500).SendString("Could not parse body data from POST request")
	}

	orderItem.UUID = uuid.NewString()
	database.DB.Create(&orderItem)
	return c.JSON(orderItem)
}

func DeleteOrderItem(c *fiber.Ctx) error {
	id := c.Params("id")

	var orderItem models.OrderItem
	database.DB.First(&orderItem, "uuid = ?", id)
	if orderItem.UUID == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error"})
	}

	database.DB.Delete(&orderItem)
	return c.Status(200).JSON(fiber.Map{"status": "successful"})
}

func UpdateOrderItem(c *fiber.Ctx) error {
	id := c.Params("id")

	var orderItem models.OrderItem
	database.DB.First(&orderItem, "uuid = ?", id)
	if orderItem.UUID == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error"})
	}

	var updatedOrderItem models.OrderItem

	err := c.BodyParser(&updatedOrderItem)
	if err != nil {
		return c.Status(503).JSON(err.Error())
	}

	database.DB.Model(&orderItem).Updates(updatedOrderItem)

	return c.Status(200).JSON(updatedOrderItem)
}
