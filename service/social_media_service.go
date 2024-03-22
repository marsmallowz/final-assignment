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
	SocialMediaService interface {
		CreateSocialMedia(ctx context.Context, req dto.SocialMediaCreateRequest, userId uint) (dto.SocialMediaResponse, error)
		GetAllSocialMedia(ctx context.Context, userId uint) (dto.GetAllSocialMediaResponse, error)
		GetSocialMediaById(ctx context.Context, socialMediaId string, userId uint) (dto.SocialMediaResponse, error)
		UpdateSocialMedia(ctx context.Context, req dto.SocialMediaUpdateRequest, socialMediaId string, userId uint) (dto.SocialMediaResponse, error)
		DeleteSocialMedia(ctx context.Context, socialMediaId string, userId uint) error
	}

	socialMediaService struct {
		socialMediaRepo repository.SocialMediaRepository
		jwtService      JWTService
	}
)

func NewSocialMediaService(socialMediaRepo repository.SocialMediaRepository, jwtService JWTService) SocialMediaService {
	return &socialMediaService{
		socialMediaRepo: socialMediaRepo,
		jwtService:      jwtService,
	}
}

func (s *socialMediaService) CreateSocialMedia(ctx context.Context, req dto.SocialMediaCreateRequest, userId uint) (dto.SocialMediaResponse, error) {

	socialMedia := entity.SocialMedia{
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
		UserID:         userId,
	}

	socialMediaPost, err := s.socialMediaRepo.CreateSocialMedia(ctx, nil, socialMedia)
	if err != nil {
		return dto.SocialMediaResponse{}, err
	}

	return dto.SocialMediaResponse{
		ID:             socialMediaPost.ID,
		Name:           socialMediaPost.Name,
		SocialMediaURL: socialMediaPost.SocialMediaURL,
		UserID:         socialMediaPost.UserID,
	}, nil
}

func (s *socialMediaService) GetAllSocialMedia(ctx context.Context, userId uint) (dto.GetAllSocialMediaResponse, error) {
	socialMedias, err := s.socialMediaRepo.GetAllSocialMedia(ctx, nil, userId)
	if err != nil {
		return dto.GetAllSocialMediaResponse{}, err
	}

	var socialMediasResponse []dto.SocialMediaResponse
	if len(socialMedias) > 0 {
		for _, socialMedia := range socialMedias {
			socialMediasResponse = append(socialMediasResponse, dto.SocialMediaResponse{
				ID:             socialMedia.ID,
				Name:           socialMedia.Name,
				SocialMediaURL: socialMedia.SocialMediaURL,
				UserID:         socialMedia.UserID,
			})
		}
	} else {
		socialMediasResponse = []dto.SocialMediaResponse{}
	}

	return dto.GetAllSocialMediaResponse{
		SocialMedias: socialMediasResponse,
	}, nil
}

func (s *socialMediaService) GetSocialMediaById(ctx context.Context, socialMediaId string, userId uint) (dto.SocialMediaResponse, error) {
	socialMedia, err := s.socialMediaRepo.GetSocialMediaById(ctx, nil, socialMediaId)
	if err != nil {
		return dto.SocialMediaResponse{}, errors.New("socialMedia not found")
	}

	if socialMedia.UserID != userId {
		return dto.SocialMediaResponse{}, errors.New("user don't have permission")
	}

	return dto.SocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
	}, nil
}

func (s *socialMediaService) UpdateSocialMedia(ctx context.Context, req dto.SocialMediaUpdateRequest, socialMediaId string, userId uint) (dto.SocialMediaResponse, error) {
	socialMedia, err := s.socialMediaRepo.GetSocialMediaById(ctx, nil, socialMediaId)
	if err != nil {
		return dto.SocialMediaResponse{}, errors.New("socialMedia not found")

	}

	if socialMedia.UserID != userId {
		return dto.SocialMediaResponse{}, errors.New("user don't have permission")
	}

	data := entity.SocialMedia{
		ID:             socialMedia.ID,
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
		UserID:         socialMedia.UserID,
	}

	socialMediaUpdate, err := s.socialMediaRepo.UpdateSocialMedia(ctx, nil, data)
	if err != nil {
		return dto.SocialMediaResponse{}, errors.New("error Update SocialMedia")
	}

	return dto.SocialMediaResponse{
		ID:             socialMediaUpdate.ID,
		Name:           socialMediaUpdate.Name,
		SocialMediaURL: socialMediaUpdate.SocialMediaURL,
		UserID:         socialMediaUpdate.UserID,
	}, nil
}

func (s *socialMediaService) DeleteSocialMedia(ctx context.Context, socialMediaId string, userId uint) error {
	socialMedia, err := s.socialMediaRepo.GetSocialMediaById(ctx, nil, socialMediaId)
	if err != nil {
		return errors.New("socialMedia not found")
	}

	if socialMedia.UserID != userId {
		return errors.New("user don't have permission")
	}

	err = s.socialMediaRepo.DeleteSocialMedia(ctx, nil, strconv.Itoa(int(socialMedia.ID)))
	if err != nil {
		return errors.New("error delete socialMedia")
	}

	return nil
}
