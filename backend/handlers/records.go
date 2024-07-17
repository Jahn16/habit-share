package handlers

import (
	"github.com/Jahn16/socialhabits/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

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
