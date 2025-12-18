package services

import (
	"algoforces/internal/domain"
	"algoforces/internal/utils"
	"context"
	"errors"
	"time"
)

type authService struct {
	userRepo domain.UserRepository
}

func NewAuthService(userRepo domain.UserRepository) domain.UserUseCase {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Signup(ctx context.Context, req *domain.SignupRequest) (*domain.AuthResponse, error) {
	existingUser, err := s.userRepo.GetByEmail(ctx, req.Email)

	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	newUser := &domain.User{
		Email:     req.Email,
		Password:  hashedPassword,
		Role:      "user",
		CreatedAt: time.Now(),
	}

	err = s.userRepo.Create(ctx, newUser)

	if err != nil {
		return nil, err
	}

	jwtToken, err := utils.GenerateToken(newUser.Id, newUser.Role, newUser.Email)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		AccessToken: jwtToken,
		User:        *newUser,
	}, nil
}

func (s *authService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.AuthResponse, error) {

	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if !utils.VerifyPassword(req.Password, user.Password) {
		return nil, errors.New("Password is incorrect")
	}

	jwtToken, err := utils.GenerateToken(user.Id, user.Role, user.Email)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		AccessToken: jwtToken,
		User:        *user,
	}, nil
}
