package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	UUID  string `json:"uuid"`
	Text  string `json:"text"`
	Title string `json:"title"`
}
