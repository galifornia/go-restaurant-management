package main

import (
	"os"

	"github.com/galifornia/go-restaurant-management/database"
	"github.com/galifornia/go-restaurant-management/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.OpenDB()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app := fiber.New()
	app.Use(logger.New())

	// Setup routes
	routes.SetupRoutes(app)

	app.Listen("localhost:" + port)
}
