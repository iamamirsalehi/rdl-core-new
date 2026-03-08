package service

import (
	"context"
	"time"

	"github.com/rdl/core/internal/modules/subscription/domain"
	"github.com/rdl/core/internal/modules/subscription/repository"
)

type subscriptionService struct {
	subRepo  repository.SubscriptionRepository
	planRepo repository.PlanRepository
}

func New(subRepo repository.SubscriptionRepository, planRepo repository.PlanRepository) SubscriptionService {
	return &subscriptionService{subRepo: subRepo, planRepo: planRepo}
}

func (s *subscriptionService) Subscribe(ctx context.Context, userID string, req *domain.CreateSubscriptionRequest) (*domain.Subscription, error) {
	plan, err := s.planRepo.FindByID(ctx, req.PlanID)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	var expiresAt time.Time
	switch plan.Interval {
	case "yearly":
		expiresAt = now.AddDate(1, 0, 0)
	default:
		expiresAt = now.AddDate(0, 1, 0)
	}

	sub := &domain.Subscription{
		UserID:    userID,
		PlanID:    req.PlanID,
		Status:    "active",
		StartedAt: now,
		ExpiresAt: expiresAt,
	}

	if err := s.subRepo.Create(ctx, sub); err != nil {
		return nil, err
	}
	return sub, nil
}

func (s *subscriptionService) GetByID(ctx context.Context, id string) (*domain.Subscription, error) {
	return s.subRepo.FindByID(ctx, id)
}

func (s *subscriptionService) GetActiveByUserID(ctx context.Context, userID string) (*domain.Subscription, error) {
	return s.subRepo.FindActiveByUserID(ctx, userID)
}

func (s *subscriptionService) List(ctx context.Context, filter domain.ListSubscriptionsFilter) ([]*domain.Subscription, error) {
	return s.subRepo.List(ctx, filter)
}

func (s *subscriptionService) Cancel(ctx context.Context, id string) error {
	sub, err := s.subRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	now := time.Now()
	sub.Status = "cancelled"
	sub.CancelledAt = &now
	return s.subRepo.Update(ctx, sub)
}

func (s *subscriptionService) ListPlans(ctx context.Context) ([]*domain.Plan, error) {
	return s.planRepo.List(ctx)
}
