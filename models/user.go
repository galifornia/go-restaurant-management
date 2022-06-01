package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserType int

const (
	ADMIN  UserType = 0
	WAITER          = 1
)

// User represents a User schema
type User struct {
	gorm.Model
	UUID     uuid.UUID `json:"_id" gorm:"primaryKey;autoIncrement:false"`
	Email    string    `json:"email" gorm:"unique"`
	Username string    `json:"username" gorm:"unique"`
	Password string    `json:"password"`
	Avatar   string    `json:"avatar"`
	Phone    string    `json:"phone"`
	Type     UserType  `json:"type"`
}

// UserErrors represent the error format for user routes
type UserErrors struct {
	Err      bool   `json:"error"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims represent the structure of the JWT token
type Claims struct {
	jwt.StandardClaims
	ID uint `gorm:"primaryKey"`
}
