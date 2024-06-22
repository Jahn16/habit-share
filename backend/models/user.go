package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Picture string  `json:"picture"`
	Habits  []Habit `json:"habits"`
}
