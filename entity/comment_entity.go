package entity

import (
	"time"
)

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Message   string `gorm:"size:200;not null"`
	UserID    uint   `gorm:"not null" json:"user_id"`
	PhotoID   uint   `gorm:"not null" json:"photo_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
