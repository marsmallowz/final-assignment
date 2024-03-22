package repository

import (
	"context"
	"final-assignment/entity"

	"gorm.io/gorm"
)

type (
	CommentRepository interface {
		PostComment(ctx context.Context, tx *gorm.DB, comment entity.Comment) (entity.Comment, error)
		GetAllComment(ctx context.Context, tx *gorm.DB) ([]entity.Comment, error)
		GetCommentById(ctx context.Context, tx *gorm.DB, commentId string) (entity.Comment, error)
		UpdateComment(ctx context.Context, tx *gorm.DB, comment entity.Comment) (entity.Comment, error)
		DeleteComment(ctx context.Context, tx *gorm.DB, commentId string) error
	}

	commentRepository struct {
		db *gorm.DB
	}
)

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) PostComment(ctx context.Context, tx *gorm.DB, comment entity.Comment) (entity.Comment, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&comment).Error; err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}

func (r *commentRepository) GetAllComment(ctx context.Context, tx *gorm.DB) ([]entity.Comment, error) {
	if tx == nil {
		tx = r.db
	}

	var comments []entity.Comment
	if err := tx.WithContext(ctx).Model(&entity.Comment{}).Find(&comments).Error; err != nil {
		return []entity.Comment{}, err
	}

	return comments, nil
}

func (r *commentRepository) GetCommentById(ctx context.Context, tx *gorm.DB, commentId string) (entity.Comment, error) {
	if tx == nil {
		tx = r.db
	}
	var comment entity.Comment
	if err := tx.WithContext(ctx).Where("id = ?", commentId).Take(&comment).Error; err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}

func (r *commentRepository) UpdateComment(ctx context.Context, tx *gorm.DB, comment entity.Comment) (entity.Comment, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&comment).Error; err != nil {
		return entity.Comment{}, err
	}

	var commentUpdate entity.Comment
	if err := tx.WithContext(ctx).First(&commentUpdate, comment.ID).Error; err != nil {
		panic(err)
	}

	return commentUpdate, nil
}

func (r *commentRepository) DeleteComment(ctx context.Context, tx *gorm.DB, commentId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Comment{}, "id = ?", commentId).Error; err != nil {
		return err
	}

	return nil
}
