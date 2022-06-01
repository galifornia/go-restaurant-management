package models

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
