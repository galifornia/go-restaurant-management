package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	UUID      string  `json:"uuid"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	FoodID    string  `json:"food_id"`
	OrderID   string  `json:"order_id"`
}
