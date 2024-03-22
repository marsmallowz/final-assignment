package seeds

import (
	"encoding/json"
	"errors"
	"final-assignment/entity"
	"io"
	"os"

	"gorm.io/gorm"
)

func ListCommentSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/comments.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)

	var listComment []entity.Comment
	if err := json.Unmarshal(jsonData, &listComment); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.Comment{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Comment{}); err != nil {
			return err
		}
	}

	for _, data := range listComment {
		var comment entity.Comment
		err := db.Where(&entity.Comment{ID: data.ID}).First(&comment).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&comment, "id = ?", data.ID).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
