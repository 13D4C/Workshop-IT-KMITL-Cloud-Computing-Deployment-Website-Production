package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Activity struct {
	ID                  uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title               string    `gorm:"size:255" json:"title"`
	Description         string    `gorm:"type:text" json:"description"`
	ImageURL            string    `gorm:"size:500" json:"image_url"`
	Location            string    `gorm:"size:255" json:"location"`
	StartDate           time.Time `json:"start_date"`
	EndDate             time.Time `json:"end_date"`
	RegisterDeadline    time.Time `json:"register_deadline"`
	MaxParticipants     int       `json:"max_participants"`
	CurrentParticipants int       `gorm:"default:0" json:"current_participants"`
	Status              string    `gorm:"size:50;default:'open'" json:"status"`
	Category            string    `gorm:"size:100" json:"category"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (a *Activity) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}
