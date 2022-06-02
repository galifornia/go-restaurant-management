package controllers

import (
	"github.com/galifornia/go-restaurant-management/database"
	"github.com/galifornia/go-restaurant-management/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetFoodsInMenu(c *fiber.Ctx) error {
	var foods []models.FoodMenu
	database.DB.Find(&foods)
	return c.JSON(foods)
}

func AddFoodToMenu(c *fiber.Ctx) error {
	var food models.FoodMenu

	if err := c.BodyParser(&food); err != nil {
		c.Status(500).SendString("Could not parse body data from POST request")
	}

	food.UUID = uuid.NewString()
	database.DB.Create(&food)
	return c.JSON(food)
}

func RemoveFoodFromMenu(c *fiber.Ctx) error {
	var foods []models.FoodMenu
	database.DB.Find(&foods)
	return c.JSON(foods)
}
