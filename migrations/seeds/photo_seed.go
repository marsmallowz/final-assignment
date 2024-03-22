package seeds

import (
	"encoding/json"
	"errors"
	"final-assignment/entity"
	"io"
	"os"

	"gorm.io/gorm"
)

func ListPhotoSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/photos.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)

	var listPhoto []entity.Photo
	if err := json.Unmarshal(jsonData, &listPhoto); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.Photo{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Photo{}); err != nil {
			return err
		}
	}

	for _, data := range listPhoto {
		var photo entity.Photo
		err := db.Where(&entity.Photo{ID: data.ID}).First(&photo).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&photo, "id = ?", data.ID).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
