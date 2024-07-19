package handlers

import (
	"fmt"

	"github.com/Jahn16/socialhabits/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HabitList(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var habits []models.Habit
		db.Preload("Records").Find(&habits)
		return c.JSON(fiber.Map{
			"success": true,
			"value":   habits,
		})

	}
}

func HabitCreate(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Locals("email").(string)
		var user models.User
		result := db.Where(&models.User{Email: email}).First(&user)
		if result.Error != nil {
			return result.Error
		}

		habit := new(models.Habit)
		if err := c.BodyParser(habit); err != nil {
			return err
		}
		habit.UserID = user.ID
		db.Create(&habit)
		return c.JSON(fiber.Map{
			"sucess": true,
			"value":  habit,
		})
	}
}

func HabitGet(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var habit models.Habit
		db.Preload("Records").First(&habit, c.Params("id"))
		return c.JSON(fiber.Map{
			"success": true,
			"value":   habit,
		})
	}
}

func DeleteHabit(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Locals("email").(string)
		var user models.User
		db.Where(&models.User{Email: email}).First(&user)
		habitID := c.Params("id")
		var habit models.Habit
		result := db.Where(&models.Habit{UserID: user.ID}).First(&habit, habitID)
		if result.Error != nil {
			return result.Error
		}
		db.Delete(&habit)
		return c.JSON(fiber.Map{
			"sucess": true,
			"value":  habit,
		})
	}
}

func UpdateHabit(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Locals("email").(string)
		var user models.User
		db.Where(&models.User{Email: email}).First(&user)
		habitID := c.Params("id")
		fmt.Println(habitID)

		var habit models.Habit
		result := db.Where(&models.Habit{UserID: user.ID}).First(&habit, habitID)
		if result.Error != nil {
			return result.Error
		}
		updatedHabit := new(models.Habit)
		if err := c.BodyParser(updatedHabit); err != nil {
			return err
		}
		habit.Icon = updatedHabit.Icon
		habit.Name = updatedHabit.Name
		habit.Goal = updatedHabit.Goal
		db.Save(&habit)
		return c.JSON(fiber.Map{
			"sucess": true,
			"value":  habit,
		})
	}
}
