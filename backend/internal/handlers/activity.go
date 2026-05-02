package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itkmitl/workshop-registration/internal/database"
	"github.com/itkmitl/workshop-registration/internal/models"
)

type ActivityHandler struct{}

func NewActivityHandler() *ActivityHandler {
	return &ActivityHandler{}
}

func (h *ActivityHandler) GetAll(c *fiber.Ctx) error {
	var activities []models.Activity

	query := database.DB.Order("start_date ASC")

	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}

	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	if search := c.Query("search"); search != "" {
		query = query.Where("title ILIKE ?", "%"+search+"%")
	}

	if err := query.Find(&activities).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "เกิดข้อผิดพลาดในการโหลดกิจกรรม",
		})
	}

	return c.JSON(fiber.Map{
		"activities": activities,
		"total":      len(activities),
	})
}

func (h *ActivityHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var activity models.Activity
	if err := database.DB.First(&activity, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "ไม่พบกิจกรรม",
		})
	}

	return c.JSON(fiber.Map{
		"activity": activity,
	})
}
