package repository

import (
	"context"
	"final-assignment/entity"

	"gorm.io/gorm"
)

type (
	PhotoRepository interface {
		PostPhoto(ctx context.Context, tx *gorm.DB, photo entity.Photo) (entity.Photo, error)
		GetAllPhoto(ctx context.Context, tx *gorm.DB) ([]entity.Photo, error)
		GetPhotoById(ctx context.Context, tx *gorm.DB, photoId string) (entity.Photo, error)
		UpdatePhoto(ctx context.Context, tx *gorm.DB, photo entity.Photo) (entity.Photo, error)
		DeletePhoto(ctx context.Context, tx *gorm.DB, photoId string) error
	}

	photoRepository struct {
		db *gorm.DB
	}
)

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{
		db: db,
	}
}

func (r *photoRepository) PostPhoto(ctx context.Context, tx *gorm.DB, photo entity.Photo) (entity.Photo, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&photo).Error; err != nil {
		return entity.Photo{}, err
	}

	return photo, nil
}

func (r *photoRepository) GetAllPhoto(ctx context.Context, tx *gorm.DB) ([]entity.Photo, error) {
	if tx == nil {
		tx = r.db
	}

	var photos []entity.Photo
	if err := tx.WithContext(ctx).Find(&photos).Error; err != nil {
		return []entity.Photo{}, err
	}

	return photos, nil
}

func (r *photoRepository) GetPhotoById(ctx context.Context, tx *gorm.DB, photoId string) (entity.Photo, error) {
	if tx == nil {
		tx = r.db
	}
	var photo entity.Photo
	if err := tx.WithContext(ctx).Where("id = ?", photoId).Take(&photo).Error; err != nil {
		return entity.Photo{}, err
	}

	return photo, nil
}

func (r *photoRepository) UpdatePhoto(ctx context.Context, tx *gorm.DB, photo entity.Photo) (entity.Photo, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&photo).Error; err != nil {
		return entity.Photo{}, err
	}

	var photoUpdate entity.Photo
	if err := tx.WithContext(ctx).First(&photoUpdate, photo.ID).Error; err != nil {
		panic(err)
	}

	return photoUpdate, nil
}

func (r *photoRepository) DeletePhoto(ctx context.Context, tx *gorm.DB, photoId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Photo{}, "id = ?", photoId).Error; err != nil {
		return err
	}

	return nil
}
