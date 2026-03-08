package repository

import (
	"context"

	"github.com/rdl/core/internal/modules/subscription/domain"
)

type SubscriptionRepository interface {
	Create(ctx context.Context, sub *domain.Subscription) error
	FindByID(ctx context.Context, id string) (*domain.Subscription, error)
	FindActiveByUserID(ctx context.Context, userID string) (*domain.Subscription, error)
	List(ctx context.Context, filter domain.ListSubscriptionsFilter) ([]*domain.Subscription, error)
	Update(ctx context.Context, sub *domain.Subscription) error
}

type PlanRepository interface {
	FindByID(ctx context.Context, id string) (*domain.Plan, error)
	List(ctx context.Context) ([]*domain.Plan, error)
}
