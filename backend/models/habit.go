package models

import (
	"time"

	"gorm.io/gorm"
)

type Habit struct {
	gorm.Model
	Icon    string        `json:"icon"`
	Name    string        `json:"name"`
	Goal    int           `json:"goal"`
	Records []HabitRecord `json:"records"`
	UserID  uint          `json:"userID"`
}

type HabitRecord struct {
	gorm.Model
	HabitID uint      `json:"habitID"`
	Date    time.Time `json:"date"`
}
