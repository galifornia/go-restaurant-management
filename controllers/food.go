package controllers

import (
	"github.com/galifornia/go-restaurant-management/database"
	"github.com/galifornia/go-restaurant-management/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllFood(c *fiber.Ctx) error {
	var foods []models.Food
	database.DB.Find(&foods)
	return c.JSON(foods)
}

func GetFoodDetails(c *fiber.Ctx) error {
	var food models.Food
	id := c.Params("id")

	database.DB.First(&food, "uuid = ?", id)
	if food.Name == "" {
		return c.Status(500).SendString("Could not find food with provided 'id'")
	}

	return c.JSON(food)
}

func NewFood(c *fiber.Ctx) error {
	var food models.Food

	if err := c.BodyParser(&food); err != nil {
		c.Status(500).SendString("Could not parse body data from POST request")
	}

	food.UUID = uuid.NewString()
	database.DB.Create(&food)
	return c.JSON(food)
}

func DeleteFood(c *fiber.Ctx) error {
	id := c.Params("id")

	var food models.Food
	database.DB.First(&food, "uuid = ?", id)
	if food.Name == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error"})
	}

	database.DB.Delete(&food)
	return c.Status(200).JSON(fiber.Map{"status": "successful"})
}

func UpdateFoodDetails(c *fiber.Ctx) error {
	id := c.Params("id")

	var food models.Food
	database.DB.First(&food, "uuid = ?", id)
	if food.Name == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error"})

	}

	var updatedFood models.Food

	err := c.BodyParser(&updatedFood)
	if err != nil {
		return c.Status(503).JSON(err.Error())
	}

	database.DB.Model(&food).Updates(updatedFood)

	return c.Status(200).JSON(updatedFood)
}
