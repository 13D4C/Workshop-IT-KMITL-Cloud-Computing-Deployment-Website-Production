package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/itkmitl/workshop-registration/internal/database"
	"github.com/itkmitl/workshop-registration/internal/models"
	"gorm.io/gorm"
)

type RegistrationHandler struct{}

func NewRegistrationHandler() *RegistrationHandler {
	return &RegistrationHandler{}
}

func (h *RegistrationHandler) Register(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var req models.RegisterActivityRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	activityID, err := uuid.Parse(req.ActivityID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid activity ID",
		})
	}

	var activity models.Activity
	if err := database.DB.First(&activity, "id = ?", activityID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "ไม่พบกิจกรรม",
		})
	}

	if activity.Status != "open" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "กิจกรรมนี้ปิดรับสมัครแล้ว",
		})
	}

	if activity.CurrentParticipants >= activity.MaxParticipants {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "กิจกรรมนี้เต็มแล้ว",
		})
	}

	var existingReg models.Registration
	if err := database.DB.Where("user_id = ? AND activity_id = ? AND status = 'registered'", userID, activityID).First(&existingReg).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "คุณสมัครกิจกรรมนี้แล้ว",
		})
	}

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		registration := models.Registration{
			UserID:     userID,
			ActivityID: activityID,
			Status:     "registered",
			Note:       req.Note,
		}

		if err := tx.Create(&registration).Error; err != nil {
			return err
		}

		if err := tx.Model(&activity).Update("current_participants", gorm.Expr("current_participants + 1")).Error; err != nil {
			return err
		}

		if activity.CurrentParticipants+1 >= activity.MaxParticipants {
			if err := tx.Model(&activity).Update("status", "full").Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "เกิดข้อผิดพลาดในการสมัครกิจกรรม",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "สมัครกิจกรรมสำเร็จ",
	})
}

func (h *RegistrationHandler) GetMyRegistrations(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var registrations []models.Registration
	if err := database.DB.Preload("Activity").Where("user_id = ? AND status = 'registered'", userID).Order("registered_at DESC").Find(&registrations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "เกิดข้อผิดพลาดในการโหลดข้อมูล",
		})
	}

	return c.JSON(fiber.Map{
		"registrations": registrations,
		"total":         len(registrations),
	})
}

func (h *RegistrationHandler) Cancel(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	regID := c.Params("id")

	var registration models.Registration
	if err := database.DB.Where("id = ? AND user_id = ? AND status = 'registered'", regID, userID).First(&registration).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "ไม่พบการสมัครกิจกรรม",
		})
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&registration).Update("status", "cancelled").Error; err != nil {
			return err
		}

		if err := tx.Model(&models.Activity{}).Where("id = ?", registration.ActivityID).
			Update("current_participants", gorm.Expr("GREATEST(current_participants - 1, 0)")).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.Activity{}).Where("id = ? AND status = 'full'", registration.ActivityID).
			Update("status", "open").Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "เกิดข้อผิดพลาดในการยกเลิก",
		})
	}

	return c.JSON(fiber.Map{
		"message": "ยกเลิกการสมัครสำเร็จ",
	})
}

func (h *RegistrationHandler) CheckRegistration(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	activityID := c.Params("activityId")

	var registration models.Registration
	err := database.DB.Where("user_id = ? AND activity_id = ? AND status = 'registered'", userID, activityID).First(&registration).Error

	if err != nil {
		return c.JSON(fiber.Map{
			"registered": false,
		})
	}

	return c.JSON(fiber.Map{
		"registered":      true,
		"registration_id": registration.ID,
	})
}
