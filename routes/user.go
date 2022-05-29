package routes

import (
	"time"

	"github.com/galifornia/go-restaurant-management/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func setupUserRoutes(router fiber.Router) {
	userApi := router.Group("/user")
	userApi.Get("/", getUser)
}

func setupUserUnrestrictedRoutes(router fiber.Router) {
	userApi := router.Group("/user")
	userApi.Post("/login", login)
}

func getUser(c *fiber.Ctx) error {
	var user models.User
	// TODO: call database
	return c.JSON(fiber.Map{"user": user})
}

func login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		HTTPOnly: true,
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 72),
	})

	return c.JSON(fiber.Map{"token": t})
}
