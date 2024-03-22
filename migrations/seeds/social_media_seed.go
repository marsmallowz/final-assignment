package seeds

import (
	"encoding/json"
	"errors"
	"final-assignment/entity"
	"io"
	"os"

	"gorm.io/gorm"
)

func ListSocialMediaSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/social_medias.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)

	var listSocialMedia []entity.SocialMedia
	if err := json.Unmarshal(jsonData, &listSocialMedia); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.SocialMedia{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.SocialMedia{}); err != nil {
			return err
		}
	}

	for _, data := range listSocialMedia {
		var socialMedia entity.SocialMedia
		err := db.Where(&entity.SocialMedia{ID: data.ID}).First(&socialMedia).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&socialMedia, "id = ?", data.ID).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
