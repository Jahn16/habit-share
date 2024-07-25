package models

import (
	"time"
)

type User struct {
	ID           string    `json:"string"`
	Name         string    `json:"name"`
	Picture      string    `json:"picture"`
	ColorPalette *string   `json:"colorPalette"`
	Habits       []Habit   `json:"habits"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
