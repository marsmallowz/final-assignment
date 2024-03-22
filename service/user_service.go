package service

import (
	"context"
	"errors"
	"final-assignment/dto"
	"final-assignment/entity"
	"final-assignment/helpers"
	"final-assignment/repository"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error)
		UpdateUser(ctx context.Context, req dto.UserUpdateRequest, userId uint) (dto.UserResponse, error)
		DeleteUser(ctx context.Context, userId uint) error
		Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
	}

	userService struct {
		userRepo   repository.UserRepository
		jwtService JWTService
	}
)

func NewUserService(userRepo repository.UserRepository, jwtService JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (s *userService) RegisterUser(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error) {
	_, flagEmail, _ := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if flagEmail {
		return dto.UserResponse{}, errors.New("email already exist")
	}

	_, flagUsername, _ := s.userRepo.CheckUsername(ctx, nil, req.Username)
	if flagUsername {
		return dto.UserResponse{}, errors.New("username already exist")
	}

	user := entity.User{
		Email:           req.Email,
		Username:        req.Username,
		Password:        req.Password,
		Age:             req.Age,
		ProfileImageURL: req.ProfileImageURL,
	}

	userReg, err := s.userRepo.RegisterUser(ctx, nil, user)
	if err != nil {
		return dto.UserResponse{}, errors.New("failed to create user")
	}

	return dto.UserResponse{
		ID:              userReg.ID,
		Email:           userReg.Email,
		Username:        userReg.Username,
		Age:             userReg.Age,
		ProfileImageURL: userReg.ProfileImageURL,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, req dto.UserUpdateRequest, userId uint) (dto.UserResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.UserResponse{}, errors.New("user not found")
	}

	data := entity.User{
		ID:              user.ID,
		Email:           req.Email,
		Username:        req.Username,
		Age:             req.Age,
		ProfileImageURL: req.ProfileImageURL,
	}

	userUpdate, err := s.userRepo.UpdateUser(ctx, nil, data)
	if err != nil {
		return dto.UserResponse{}, errors.New("failed to update user")
	}

	return dto.UserResponse{
		ID:              userUpdate.ID,
		Email:           userUpdate.Email,
		Username:        userUpdate.Username,
		Age:             userUpdate.Age,
		ProfileImageURL: req.ProfileImageURL,
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, userId uint) error {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return errors.New("user not found")
	}
	err = s.userRepo.DeleteUser(ctx, nil, user.ID)
	if err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}

func (s *userService) Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	check, flag, err := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if err != nil || !flag {
		return dto.UserLoginResponse{}, errors.New("email not found")
	}

	checkPassword, err := helpers.CheckPassword(check.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.UserLoginResponse{}, errors.New("password not match")
	}

	token := s.jwtService.GenerateToken(check.ID)

	return dto.UserLoginResponse{
		Token: token,
	}, nil
}
