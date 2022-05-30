package helpers

import (
	"time"

	"github.com/galifornia/go-restaurant-management/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret") // []byte(os.Getenv("PRIV_KEY")) // !FIXME

// GenerateTokens generates the access and refresh tokens
func GenerateToken(uuid string) string {
	accessToken := GenerateAccessClaims(uuid)

	return accessToken
}

// GenerateAccessClaims returns a claim and a access_token string
func GenerateAccessClaims(uuid string) string {
	t := time.Now()
	claim := &models.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    uuid,
			ExpiresAt: t.Add(10 * time.Hour).Unix(),
			Subject:   "access_token",
			IssuedAt:  t.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}

func GetAuthCookie(accessToken string) *fiber.Cookie {
	accessCookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(10 * 24 * time.Hour),
		HTTPOnly: true,
	}

	return accessCookie
}

// IsEmpty checks if a string is empty
func IsEmpty(str string) (bool, string) {
	if str == "" {
		return true, "Must not be empty"
	}

	return false, ""
}

// ValidateRegister: validate creadentials
func ValidateRegister(u *models.User) *models.UserErrors {
	e := &models.UserErrors{Err: false}
	e.Err, e.Username = IsEmpty(u.Username)

	// TODO: validate email & username

	return e
}
