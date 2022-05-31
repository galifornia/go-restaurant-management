package main

import (
	"os"

	"github.com/galifornia/go-restaurant-management/config"
	"github.com/galifornia/go-restaurant-management/database"
	"github.com/galifornia/go-restaurant-management/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadConfig()
	database.OpenDB()

	port := os.Getenv("PORT")

	app := fiber.New()
	app.Use(logger.New())

	// Setup routes
	routes.SetupRoutes(app)

	app.Listen("localhost:" + port)
}
