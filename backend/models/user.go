package models

import (
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Picture      string    `json:"picture"`
	ColorPalette *string   `json:"colorPalette"`
	Habits       []Habit   `json:"habits"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Friends      *[]User   `gorm:"many2many:user_friends" json:"friends"`
}
