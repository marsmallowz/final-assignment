package service

import (
	"context"
	"errors"
	"final-assignment/dto"
	"final-assignment/entity"
	"final-assignment/repository"
	"strconv"
)

type (
	CommentService interface {
		PostComment(ctx context.Context, req dto.CommentCreateRequest, userId uint) (dto.CommentResponse, error)
		GetAllComment(ctx context.Context) (dto.GetAllCommentResponse, error)
		GetCommentById(ctx context.Context, commentId string) (dto.CommentResponse, error)
		UpdateComment(ctx context.Context, req dto.CommentUpdateRequest, commentId string, userId uint) (dto.CommentResponse, error)
		DeleteComment(ctx context.Context, commentId string, userId uint) error
	}

	commentService struct {
		commentRepo repository.CommentRepository
		jwtService  JWTService
	}
)

func NewCommentService(commentRepo repository.CommentRepository, jwtService JWTService) CommentService {
	return &commentService{
		commentRepo: commentRepo,
		jwtService:  jwtService,
	}
}

func (s *commentService) PostComment(ctx context.Context, req dto.CommentCreateRequest, userId uint) (dto.CommentResponse, error) {

	comment := entity.Comment{
		Message: req.Message,
		PhotoID: req.PhotoID,
		UserID:  userId,
	}

	commentPost, err := s.commentRepo.PostComment(ctx, nil, comment)
	if err != nil {
		return dto.CommentResponse{}, err
	}

	return dto.CommentResponse{
		ID:      commentPost.ID,
		Message: commentPost.Message,
		PhotoID: commentPost.PhotoID,
		UserID:  commentPost.UserID,
	}, nil
}

func (s *commentService) GetAllComment(ctx context.Context) (dto.GetAllCommentResponse, error) {
	comments, err := s.commentRepo.GetAllComment(ctx, nil)
	if err != nil {
		return dto.GetAllCommentResponse{}, err
	}

	var commentsResponse []dto.CommentResponse
	if len(comments) > 0 {
		for _, comment := range comments {
			commentsResponse = append(commentsResponse, dto.CommentResponse{
				ID:      comment.ID,
				Message: comment.Message,
				PhotoID: comment.PhotoID,
				UserID:  comment.UserID,
			})
		}
	} else {
		commentsResponse = []dto.CommentResponse{}
	}

	return dto.GetAllCommentResponse{
		Comments: commentsResponse,
	}, nil
}

func (s *commentService) GetCommentById(ctx context.Context, commentId string) (dto.CommentResponse, error) {
	comment, err := s.commentRepo.GetCommentById(ctx, nil, commentId)
	if err != nil {
		return dto.CommentResponse{}, errors.New("comment not found")

	}

	return dto.CommentResponse{
		ID:      comment.ID,
		Message: comment.Message,
		PhotoID: comment.PhotoID,
		UserID:  comment.UserID,
	}, nil
}

func (s *commentService) UpdateComment(ctx context.Context, req dto.CommentUpdateRequest, commentId string, userId uint) (dto.CommentResponse, error) {
	comment, err := s.commentRepo.GetCommentById(ctx, nil, commentId)
	if err != nil {
		return dto.CommentResponse{}, errors.New("comment not found")

	}

	if comment.UserID != userId {
		return dto.CommentResponse{}, errors.New("user don't have permission")
	}

	data := entity.Comment{
		ID:      comment.ID,
		Message: req.Message,
		UserID:  comment.UserID,
	}

	commentUpdate, err := s.commentRepo.UpdateComment(ctx, nil, data)
	if err != nil {
		return dto.CommentResponse{}, errors.New("error Update Comment")
	}

	return dto.CommentResponse{
		ID:      commentUpdate.ID,
		Message: commentUpdate.Message,
		PhotoID: commentUpdate.PhotoID,
		UserID:  commentUpdate.UserID,
	}, nil
}

func (s *commentService) DeleteComment(ctx context.Context, commentId string, userId uint) error {
	comment, err := s.commentRepo.GetCommentById(ctx, nil, commentId)
	if err != nil {
		return errors.New("comment not found")
	}

	if comment.UserID != userId {
		return errors.New("user don't have permission")
	}

	err = s.commentRepo.DeleteComment(ctx, nil, strconv.Itoa(int(comment.ID)))
	if err != nil {
		return errors.New("error delete comment")
	}

	return nil
}
