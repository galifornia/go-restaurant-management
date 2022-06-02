package controllers

import (
	"github.com/galifornia/go-restaurant-management/database"
	"github.com/galifornia/go-restaurant-management/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllNotes(c *fiber.Ctx) error {
	var notes []models.Note
	database.DB.Find(&notes)
	return c.JSON(notes)
}

func GetNote(c *fiber.Ctx) error {
	var note models.Note
	id := c.Params("id")

	database.DB.First(&note, "uuid = ?", id)
	if note.UUID == "" {
		return c.Status(500).SendString("Could not find Note with provided 'id'")
	}

	return c.JSON(note)
}

func NewNote(c *fiber.Ctx) error {
	var note models.Note

	if err := c.BodyParser(&note); err != nil {
		c.Status(500).SendString("Could not parse body data from POST request")
	}

	note.UUID = uuid.NewString()
	database.DB.Create(&note)
	return c.JSON(note)
}

func DeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")

	var note models.Note
	database.DB.First(&note, "uuid = ?", id)
	if note.UUID == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error"})
	}

	database.DB.Delete(&note)
	return c.Status(200).JSON(fiber.Map{"status": "successful"})
}

func UpdateNote(c *fiber.Ctx) error {
	id := c.Params("id")

	var note models.Note
	database.DB.First(&note, "uuid = ?", id)
	if note.UUID == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error"})
	}

	var updatedNote models.Note

	err := c.BodyParser(&updatedNote)
	if err != nil {
		return c.Status(503).JSON(err.Error())
	}

	database.DB.Model(&note).Updates(updatedNote)

	return c.Status(200).JSON(updatedNote)
}
