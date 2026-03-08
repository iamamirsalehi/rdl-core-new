package service

import (
	"context"

	"github.com/rdl/core/internal/modules/subscription/domain"
)

type SubscriptionService interface {
	Subscribe(ctx context.Context, userID string, req *domain.CreateSubscriptionRequest) (*domain.Subscription, error)
	GetByID(ctx context.Context, id string) (*domain.Subscription, error)
	GetActiveByUserID(ctx context.Context, userID string) (*domain.Subscription, error)
	List(ctx context.Context, filter domain.ListSubscriptionsFilter) ([]*domain.Subscription, error)
	Cancel(ctx context.Context, id string) error
	ListPlans(ctx context.Context) ([]*domain.Plan, error)
}
