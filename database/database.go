package database

import "github.com/Jahn16/habitshare/models"

var db []*models.Habit

func Setup() {
	db = make([]*models.Habit, 0)
}

func Insert(habit *models.Habit) {
	db = append(db, habit)
}

func List() []*models.Habit {
	return db
}
