package service

import (
	"context"
	"errors"

	"github.com/rdl/core/internal/modules/auth/domain"
	"github.com/rdl/core/internal/modules/auth/repository"
)

type authService struct {
	userRepo repository.UserRepository
	// tokenManager TokenManager (inject JWT or similar)
}

func New(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(ctx context.Context, req *domain.RegisterRequest) (*domain.User, error) {
	existing, _ := s.userRepo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	user := &domain.User{
		Name:  req.Name,
		Email: req.Email,
		Role:  "user",
		// Password should be hashed before storing
		Password: req.Password,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.TokenPair, error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// TODO: verify hashed password
	_ = user

	return &domain.TokenPair{
		AccessToken:  "access_token_placeholder",
		RefreshToken: "refresh_token_placeholder",
	}, nil
}

func (s *authService) Refresh(ctx context.Context, req *domain.RefreshRequest) (*domain.TokenPair, error) {
	// TODO: validate refresh token and issue new pair
	return &domain.TokenPair{}, nil
}

func (s *authService) Logout(ctx context.Context, userID string) error {
	// TODO: invalidate token in redis
	return nil
}

func (s *authService) ValidateToken(ctx context.Context, token string) (*domain.User, error) {
	// TODO: parse and validate JWT
	return nil, nil
}
