package db

import (
	"go-server-start/internal/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {

	db, err := gorm.Open(sqlite.Open(config.DatabasePath), &gorm.Config{})
	if err != nil {
		return err
	}

	// SkipDefaultTransaction is used to skip the default transaction
	db.SkipDefaultTransaction = true

	DB = db

	registerModels()

	return nil
}
