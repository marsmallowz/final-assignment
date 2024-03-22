package entity

import (
	"time"
)

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null;size:50"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url"`
	UserID         uint   `gorm:"not null" json:"user_id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
