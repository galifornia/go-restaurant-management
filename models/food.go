package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	UUID  string   `json:"uuid"`
	Name  *string  `json:"name"`
	Price *float64 `json:"price"`
	Img   *string  `json:"img"`
	// MenuID string `json:"menu_id"``
}
