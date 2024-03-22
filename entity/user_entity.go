package entity

import (
	"final-assignment/helpers"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint   `gorm:"primaryKey"`
	Username        string `gorm:"size:50;uniqueIndex;not null;"`
	Password        string `gorm:"not null"`
	Email           string `gorm:"size:150;uniqueIndex;not null"`
	Age             uint8  `gorm:"not null"`
	ProfileImageURL string `json:"profile_image_url"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Photos          []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments        []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SocialMedia     []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
