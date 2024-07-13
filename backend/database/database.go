package database

import (
	"github.com/Jahn16/socialhabits/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.User{}, &models.Habit{}, &models.HabitRecord{})
	return db, err
}
