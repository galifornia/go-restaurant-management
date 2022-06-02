package models

import "gorm.io/gorm"

type FoodMenu struct {
	gorm.Model
	UUID   string `json:"uuid"`
	MenuID string `json:"menu-id"`
	FoodID string `json:"food-id"`
}
