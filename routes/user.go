package routes

import (
	"github.com/galifornia/go-restaurant-management/database"
	"github.com/galifornia/go-restaurant-management/helpers"
	"github.com/galifornia/go-restaurant-management/middleware"
	"github.com/galifornia/go-restaurant-management/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func setupUserRoutes(router fiber.Router) {
	userInfoApi := router.Group("/userinfo")
	userInfoApi.Use(middleware.SecureAuth())
	userInfoApi.Get("/", getUser)

	userApi := router.Group("/user")
	userApi.Post("/signup", signUp)
	userApi.Post("/login", logIn)
}

// LoginUser route logins a user in the app
func logIn(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}

	input := new(LoginInput)

	if err := c.BodyParser(input); err != nil {
		return c.JSON(fiber.Map{"error": true, "input": "Please review your input"})
	}

	if input.Identity == "" {
		return c.JSON(fiber.Map{"error": true, "input": "Please review your input"})
	}
	user := new(models.User)

	if res := database.DB.Where(
		&models.User{Email: input.Identity}).Or(
		&models.User{Username: input.Identity},
	).First(&user); res.RowsAffected <= 0 {
		return c.JSON(fiber.Map{"error": true, "general": "Invalid Credentials."})
	}

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.JSON(fiber.Map{"error": true, "general": "Invalid Credentials."})
	}

	// setting up the authorization cookies

	accessToken := helpers.GenerateToken(user.UUID.String())
	accessCookie := helpers.GetAuthCookie(accessToken)

	c.Cookie(accessCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": accessToken,
	})
}

func signUp(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.JSON(fiber.Map{
			"error": true,
			"input": "Please review your input",
		})
	}

	// validate if the email, username and password are in correct format
	errors := helpers.ValidateRegister(user)
	if errors.Err {
		return c.JSON(errors)
	}

	if count := database.DB.Where(&models.User{Email: user.Email}).First(new(models.User)).RowsAffected; count > 0 {
		errors.Err, errors.Email = true, "Email is already registered"
	}
	if count := database.DB.Where(&models.User{Username: user.Username}).First(new(models.User)).RowsAffected; count > 0 {
		errors.Err, errors.Username = true, "Username is already registered"
	}
	if errors.Err {
		return c.JSON(errors)
	}

	user.UUID = uuid.New()

	// Hashing the password with a random salt
	password := []byte(user.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(
		password,
		10,
	)

	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		return c.JSON(fiber.Map{
			"error":   true,
			"general": "Something went wrong, please try again later. 😕",
		})
	}

	// setting up the authorization cookies
	accessToken := helpers.GenerateToken(user.UUID.String())
	accessCookie := helpers.GetAuthCookie(accessToken)
	c.Cookie(accessCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": accessToken,
	})
}

func getUser(c *fiber.Ctx) error {
	var user models.User
	id := c.Locals("id")
	database.DB.First(&user, "uuid = ?", id)

	if user.Username == "" {
		return c.Status(500).SendString("No user found with that ID. Please logIn again.")
	}

	return c.JSON(fiber.Map{"user": user})
}
