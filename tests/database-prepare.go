package main

import (
	"go-server-start/internal/config"
	"go-server-start/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open(config.AppConfig.Database.DBName), &gorm.Config{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	db.Create([]*models.User{
		{
			ID:   1,
			Name: "John",
		},
		{
			ID:   2,
			Name: "Doe",
		},
		{
			ID:   3,
			Name: "Smith",
		},
		{
			ID:   4,
			Name: "Jane",
		},
		{
			ID:   5,
			Name: "Peter",
		},
		{
			ID:   6,
			Name: "James",
		},
	})

}
