package entity

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Image      string         `json:"image"`
	CategoryID uint           `json:"category_id"`
	CreatedAt  time.Time      `json:"creted_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
