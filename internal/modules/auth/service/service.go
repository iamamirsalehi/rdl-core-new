package service

import (
	"context"

	"github.com/rdl/core/internal/modules/auth/domain"
)

type AuthService interface {
	Register(ctx context.Context, req *domain.RegisterRequest) (*domain.User, error)
	Login(ctx context.Context, req *domain.LoginRequest) (*domain.TokenPair, error)
	Refresh(ctx context.Context, req *domain.RefreshRequest) (*domain.TokenPair, error)
	Logout(ctx context.Context, userID string) error
	ValidateToken(ctx context.Context, token string) (*domain.User, error)
}
