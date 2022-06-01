package routes

import (
	"github.com/galifornia/go-restaurant-management/controllers"
	"github.com/galifornia/go-restaurant-management/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupUserRoutes(router fiber.Router) {
	userInfoApi := router.Group("/users")
	userInfoApi.Use(middleware.SecureAuth())
	userInfoApi.Get("/", controllers.GetAllUsers)
	userInfoApi.Get("/", controllers.GetUser)

	userApi := router.Group("/user")
	userApi.Post("/signup", controllers.SignUp)
	userApi.Post("/login", controllers.LogIn)
}
