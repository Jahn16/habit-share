package database

import (
	"github.com/Jahn16/habitshare/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Setup() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	db.AutoMigrate(&models.Habit{})
	return db, err
}
