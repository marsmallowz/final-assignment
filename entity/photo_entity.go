package entity

import (
	"time"
)

type Photo struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"size:100;not null"`
	Caption   string    `gorm:"size:200"`
	PhotoUrl  string    `gorm:"not null" json:"photo_url"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Comments  []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
