package database

import (
	"github.com/galifornia/go-restaurant-management/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDB() *gorm.DB {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("Unable to open a database connection")
	}

	// Migrate the schema
	DB.AutoMigrate(&models.User{}, &models.Food{}, &models.Invoice{}, &models.Order{}, &models.OrderItem{}, &models.Table{}, &models.Menu{}, &models.FoodMenu{})

	return DB
}
