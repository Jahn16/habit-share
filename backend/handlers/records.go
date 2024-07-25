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
		userID := c.Locals("id").(string)
		habitID := c.Params("id")
		var habit models.Habit
		result := db.Where(&models.Habit{UserID: userID}).First(&habit, habitID)
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

		userID := c.Locals("id").(string)
		var habit models.Habit
		habitResult := db.Where(&models.Habit{UserID: userID}).First(&habit, record.HabitID)
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
