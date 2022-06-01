package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UUID    string    `json:"uuid"`
	TableID *string   `json:"table_id"`
	Date    time.Time `json:"date"`
}
