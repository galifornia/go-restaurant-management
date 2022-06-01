package models

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	UUID   string `json:"uuid"`
	Chairs int    `json:"chairs"`
}
