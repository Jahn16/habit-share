package handlers

import (
	"github.com/Jahn16/habitshare/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAuthenticatedUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Locals("email").(string)
		var user models.User
		db.Where(&models.User{Email: email}).Preload("Habits.Records").First(&user)
		return c.JSON(fiber.Map{
			"success": true,
			"value":   user,
		})
	}
}

func UserCreate(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Locals("username").(string)
		email := c.Locals("email").(string)
		picture := c.Locals("picture").(string)
		user := models.User{
			Name:    username,
			Email:   email,
			Picture: picture,
		}
		db.Create(&user)
		return c.JSON(fiber.Map{
			"sucess": true,
			"value":  user,
		})

	}
}

func UserGet(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.User
		result := db.Preload("Habits.Records").First(&user, c.Params("id"))
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"sucess": "false", "message": "User not found"})
		}
		return c.JSON(fiber.Map{
			"success": true,
			"value":   user,
		})
	}
}

func UserList(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var users []models.User
		db.Find(&users)
		return c.JSON(fiber.Map{
			"success": true,
			"value":   users,
		})

	}
}

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

func RecordHabit(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		habitRecord := new(models.HabitRecord)
		if err := c.BodyParser(habitRecord); err != nil {
			return err
		}
		email := c.Locals("email").(string)
		var user models.User
		db.Where(&models.User{Email: email}).First(&user)
		var habit models.Habit
		habitID := c.Params("id")
		result := db.Where(&models.Habit{UserID: user.ID}).First(&habit, habitID)
		if result.Error != nil {
			return result.Error
		}
		habitRecord.HabitID = habit.ID
		db.Create(&habitRecord)
		return c.JSON(fiber.Map{
			"sucess": true,
			"value":  habitRecord,
		})
	}
}

func DeleteRecord(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		recordID := c.Params("id")
		var record models.HabitRecord
		recordResult := db.First(&record, recordID)
		if recordResult.Error != nil {
			return recordResult.Error
		}

		email := c.Locals("email").(string)
		var user models.User
		db.Where(&models.User{Email: email}).First(&user)
		var habit models.Habit
		habitResult := db.Where(&models.Habit{UserID: user.ID}).First(&habit, record.HabitID)
		if habitResult.Error != nil {
			return habitResult.Error
		}
		db.Delete(&record)
		return c.JSON(fiber.Map{
			"sucess": true,
			"value":  record,
		})
	}
}
