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
	PhotoService interface {
		PostPhoto(ctx context.Context, req dto.PhotoCreateRequest, userId uint) (dto.PhotoResponse, error)
		GetPhotoById(ctx context.Context, photoId string) (dto.PhotoResponse, error)
		GetAllPhoto(ctx context.Context) (dto.GetAllPhotoResponse, error)
		UpdatePhoto(ctx context.Context, req dto.PhotoUpdateRequest, photoId string, userId uint) (dto.PhotoResponse, error)
		DeletePhoto(ctx context.Context, photoId string, userId uint) error
	}

	photoService struct {
		photoRepo  repository.PhotoRepository
		jwtService JWTService
	}
)

func NewPhotoService(photoRepo repository.PhotoRepository, jwtService JWTService) PhotoService {
	return &photoService{
		photoRepo:  photoRepo,
		jwtService: jwtService,
	}
}

func (s *photoService) PostPhoto(ctx context.Context, req dto.PhotoCreateRequest, userId uint) (dto.PhotoResponse, error) {

	photo := entity.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoURL,
		UserID:   userId,
	}

	photoPost, err := s.photoRepo.PostPhoto(ctx, nil, photo)
	if err != nil {
		return dto.PhotoResponse{}, err
	}

	return dto.PhotoResponse{
		ID:       photoPost.ID,
		Title:    photoPost.Title,
		Caption:  photoPost.Caption,
		PhotoURL: photoPost.PhotoUrl,
		UserID:   photoPost.UserID,
	}, nil
}

func (s *photoService) GetAllPhoto(ctx context.Context) (dto.GetAllPhotoResponse, error) {
	photos, err := s.photoRepo.GetAllPhoto(ctx, nil)
	if err != nil {
		return dto.GetAllPhotoResponse{}, err
	}
	var photosResponse []dto.PhotoResponse
	if len(photos) > 0 {
		for _, photo := range photos {
			photosResponse = append(photosResponse, dto.PhotoResponse{
				ID:       photo.ID,
				Title:    photo.Title,
				Caption:  photo.Caption,
				PhotoURL: photo.PhotoUrl,
				UserID:   photo.UserID,
			})
		}
	} else {
		photosResponse = []dto.PhotoResponse{}
	}

	return dto.GetAllPhotoResponse{
		Photos: photosResponse,
	}, nil
}

func (s *photoService) GetPhotoById(ctx context.Context, photoId string) (dto.PhotoResponse, error) {
	photo, err := s.photoRepo.GetPhotoById(ctx, nil, photoId)
	if err != nil {
		return dto.PhotoResponse{}, errors.New("photo not found")

	}

	return dto.PhotoResponse{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoUrl,
		UserID:   photo.UserID,
	}, nil
}

func (s *photoService) UpdatePhoto(ctx context.Context, req dto.PhotoUpdateRequest, photoId string, userId uint) (dto.PhotoResponse, error) {
	photo, err := s.photoRepo.GetPhotoById(ctx, nil, photoId)
	if err != nil {
		return dto.PhotoResponse{}, errors.New("photo not found")

	}

	if photo.UserID != userId {
		return dto.PhotoResponse{}, errors.New("user don't have permission")
	}

	data := entity.Photo{
		ID:       photo.ID,
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoURL,
		UserID:   photo.UserID,
	}

	photoUpdate, err := s.photoRepo.UpdatePhoto(ctx, nil, data)
	if err != nil {
		return dto.PhotoResponse{}, errors.New("error Update Photo")
	}

	return dto.PhotoResponse{
		ID:       photoUpdate.ID,
		Title:    photoUpdate.Title,
		Caption:  photoUpdate.Caption,
		PhotoURL: photoUpdate.PhotoUrl,
		UserID:   photoUpdate.UserID,
	}, nil
}

func (s *photoService) DeletePhoto(ctx context.Context, photoId string, userId uint) error {
	photo, err := s.photoRepo.GetPhotoById(ctx, nil, photoId)
	if err != nil {
		return errors.New("photo not found")
	}

	if photo.UserID != userId {
		return errors.New("user don't have permission")
	}

	err = s.photoRepo.DeletePhoto(ctx, nil, strconv.Itoa(int(photo.ID)))
	if err != nil {
		return errors.New("error delete photo")
	}

	return nil
}
