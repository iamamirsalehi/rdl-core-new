package repository

import (
	"context"

	"github.com/rdl/core/internal/modules/cathub/domain"
)

type HubRepository interface {
	Create(ctx context.Context, hub *domain.Hub) error
	FindByID(ctx context.Context, id string) (*domain.Hub, error)
	List(ctx context.Context, filter domain.ListHubsFilter) ([]*domain.Hub, error)
	Update(ctx context.Context, hub *domain.Hub) error
	Delete(ctx context.Context, id string) error
}

type CategoryRepository interface {
	Create(ctx context.Context, category *domain.Category) error
	FindByID(ctx context.Context, id string) (*domain.Category, error)
	FindBySlug(ctx context.Context, slug string) (*domain.Category, error)
	List(ctx context.Context, parentID string) ([]*domain.Category, error)
	Update(ctx context.Context, category *domain.Category) error
	Delete(ctx context.Context, id string) error
}
