package db

import (
	"go-server-start/models"
	"gorm.io/gorm"
)

var (
	GetTableUser func() *gorm.DB
)

func registerModels() {
	GetTableUser = func() *gorm.DB { return DB.Model(&models.User{}) }
}
