package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	StudentID  string    `gorm:"uniqueIndex;size:20" json:"student_id"`
	Email      string    `gorm:"uniqueIndex;size:255" json:"email"`
	Password   string    `gorm:"size:255" json:"-"`
	FirstName  string    `gorm:"size:100" json:"first_name"`
	LastName   string    `gorm:"size:100" json:"last_name"`
	Phone      string    `gorm:"size:20" json:"phone"`
	Faculty    string    `gorm:"size:100" json:"faculty"`
	Department string    `gorm:"size:100" json:"department"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

type RegisterRequest struct {
	StudentID  string `json:"student_id" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6"`
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	Phone      string `json:"phone"`
	Faculty    string `json:"faculty"`
	Department string `json:"department"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID         uuid.UUID `json:"id"`
	StudentID  string    `json:"student_id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Phone      string    `json:"phone"`
	Faculty    string    `json:"faculty"`
	Department string    `json:"department"`
	CreatedAt  time.Time `json:"created_at"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:         u.ID,
		StudentID:  u.StudentID,
		Email:      u.Email,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Phone:      u.Phone,
		Faculty:    u.Faculty,
		Department: u.Department,
		CreatedAt:  u.CreatedAt,
	}
}
