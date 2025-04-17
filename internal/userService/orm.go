package userService

import (
	"pet_project_1_etap/internal/taskService"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint               `json:"id" gorm:"primaryKey"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeletedAt gorm.DeletedAt     `json:"deleted_at" gorm:"index"`
	Tasks     []taskService.Task `gorm:"foreignKey:UserID"`
}
