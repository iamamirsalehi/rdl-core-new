package service

import (
	"context"

	"github.com/rdl/core/internal/modules/cathub/domain"
)

type CathubService interface {
	CreateHub(ctx context.Context, ownerID string, req *domain.CreateHubRequest) (*domain.Hub, error)
	GetHubByID(ctx context.Context, id string) (*domain.Hub, error)
	ListHubs(ctx context.Context, filter domain.ListHubsFilter) ([]*domain.Hub, error)
	UpdateHub(ctx context.Context, id string, req *domain.UpdateHubRequest) (*domain.Hub, error)
	DeleteHub(ctx context.Context, id string) error
}
