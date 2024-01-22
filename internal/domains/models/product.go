package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"type:VARCHAR(255);" json:"title" validate:"required"`
	Description string         `gorm:"type:TEXT;" json:"description" validate:"required"`
	Price       float64        `gorm:"type:DECIMAL(20,2);" json:"price" validate:"required,gt=0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
