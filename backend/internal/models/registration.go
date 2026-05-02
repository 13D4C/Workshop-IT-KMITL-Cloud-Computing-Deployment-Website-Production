package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Registration struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID       uuid.UUID `gorm:"type:uuid;index" json:"user_id"`
	ActivityID   uuid.UUID `gorm:"type:uuid;index" json:"activity_id"`
	Status       string    `gorm:"size:50;default:'registered'" json:"status"`
	Note         string    `gorm:"type:text" json:"note"`
	RegisteredAt time.Time `json:"registered_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	User     User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Activity Activity `gorm:"foreignKey:ActivityID" json:"activity,omitempty"`
}

func (r *Registration) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	if r.RegisteredAt.IsZero() {
		r.RegisteredAt = time.Now()
	}
	return nil
}

type RegisterActivityRequest struct {
	ActivityID string `json:"activity_id" validate:"required"`
	Note       string `json:"note"`
}

type RegistrationResponse struct {
	ID           uuid.UUID `json:"id"`
	ActivityID   uuid.UUID `json:"activity_id"`
	Status       string    `json:"status"`
	Note         string    `json:"note"`
	RegisteredAt time.Time `json:"registered_at"`
	Activity     Activity  `json:"activity"`
}
