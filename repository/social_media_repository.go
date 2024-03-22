package repository

import (
	"context"
	"final-assignment/entity"

	"gorm.io/gorm"
)

type (
	SocialMediaRepository interface {
		CreateSocialMedia(ctx context.Context, tx *gorm.DB, socialMedia entity.SocialMedia) (entity.SocialMedia, error)
		GetAllSocialMedia(ctx context.Context, tx *gorm.DB, userId uint) ([]entity.SocialMedia, error)
		GetSocialMediaById(ctx context.Context, tx *gorm.DB, socialMediaId string) (entity.SocialMedia, error)
		UpdateSocialMedia(ctx context.Context, tx *gorm.DB, socialMedia entity.SocialMedia) (entity.SocialMedia, error)
		DeleteSocialMedia(ctx context.Context, tx *gorm.DB, socialMediaId string) error
	}

	socialMediaRepository struct {
		db *gorm.DB
	}
)

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{
		db: db,
	}
}

func (r *socialMediaRepository) CreateSocialMedia(ctx context.Context, tx *gorm.DB, socialMedia entity.SocialMedia) (entity.SocialMedia, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&socialMedia).Error; err != nil {
		return entity.SocialMedia{}, err
	}

	return socialMedia, nil
}

func (r *socialMediaRepository) GetAllSocialMedia(ctx context.Context, tx *gorm.DB, userId uint) ([]entity.SocialMedia, error) {
	if tx == nil {
		tx = r.db
	}

	var socialMedias []entity.SocialMedia
	if err := tx.WithContext(ctx).Model(&entity.SocialMedia{}).Where("user_id=?", userId).Find(&socialMedias).Error; err != nil {
		return []entity.SocialMedia{}, err
	}

	return socialMedias, nil
}

func (r *socialMediaRepository) GetSocialMediaById(ctx context.Context, tx *gorm.DB, socialMediaId string) (entity.SocialMedia, error) {
	if tx == nil {
		tx = r.db
	}
	var socialMedia entity.SocialMedia
	if err := tx.WithContext(ctx).Where("id = ?", socialMediaId).Take(&socialMedia).Error; err != nil {
		return entity.SocialMedia{}, err
	}

	return socialMedia, nil
}

func (r *socialMediaRepository) UpdateSocialMedia(ctx context.Context, tx *gorm.DB, socialMedia entity.SocialMedia) (entity.SocialMedia, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&socialMedia).Error; err != nil {
		return entity.SocialMedia{}, err
	}

	var socialMediaUpdate entity.SocialMedia
	if err := tx.WithContext(ctx).First(&socialMediaUpdate, socialMedia.ID).Error; err != nil {
		panic(err)
	}

	return socialMediaUpdate, nil
}

func (r *socialMediaRepository) DeleteSocialMedia(ctx context.Context, tx *gorm.DB, socialMediaId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.SocialMedia{}, "id = ?", socialMediaId).Error; err != nil {
		return err
	}

	return nil
}
