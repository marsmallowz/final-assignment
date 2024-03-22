package migrations

import (
	"final-assignment/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Photo{},
		&entity.Comment{},
		&entity.SocialMedia{},
	); err != nil {
		return err
	}

	return nil
}
