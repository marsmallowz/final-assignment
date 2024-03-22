package migrations

import (
	"final-assignment/migrations/seeds"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.ListUserSeeder(db); err != nil {
		return err
	}
	if err := seeds.ListSocialMediaSeeder(db); err != nil {
		return err
	}
	if err := seeds.ListPhotoSeeder(db); err != nil {
		return err
	}
	if err := seeds.ListCommentSeeder(db); err != nil {
		return err
	}

	return nil
}
