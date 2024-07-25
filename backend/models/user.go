package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string    `json:"string"`
	Name      string    `json:"name"`
	Picture   string    `json:"picture"`
	Habits    []Habit   `json:"habits"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
