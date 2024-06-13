package models

import "gorm.io/gorm"

type Habit struct {
	gorm.Model
	Name string `json:"name"`
	Goal int    `json:"goal"`
}
