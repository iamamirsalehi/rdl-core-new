package service

import (
	"context"

	"github.com/rdl/core/internal/modules/payment/domain"
	"github.com/rdl/core/internal/modules/payment/repository"
)

type paymentService struct {
	repo repository.PaymentRepository
}

func New(repo repository.PaymentRepository) PaymentService {
	return &paymentService{repo: repo}
}

func (s *paymentService) InitiatePayment(ctx context.Context, userID string, req *domain.CreatePaymentRequest) (*domain.Payment, error) {
	p := &domain.Payment{
		UserID:   userID,
		Amount:   req.Amount,
		Currency: req.Currency,
		Provider: req.Provider,
		Status:   "pending",
	}

	// TODO: call provider SDK to create charge session

	if err := s.repo.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *paymentService) GetByID(ctx context.Context, id string) (*domain.Payment, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *paymentService) List(ctx context.Context, filter domain.ListPaymentsFilter) ([]*domain.Payment, error) {
	return s.repo.List(ctx, filter)
}

func (s *paymentService) HandleWebhook(ctx context.Context, req *domain.PaymentWebhookRequest) error {
	// TODO: verify webhook signature and update payment status
	return nil
}
