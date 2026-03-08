package repository

import (
	"context"

	"github.com/rdl/core/internal/modules/payment/domain"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *domain.Payment) error
	FindByID(ctx context.Context, id string) (*domain.Payment, error)
	FindByProviderID(ctx context.Context, providerID string) (*domain.Payment, error)
	List(ctx context.Context, filter domain.ListPaymentsFilter) ([]*domain.Payment, error)
	UpdateStatus(ctx context.Context, id, status string) error
}
