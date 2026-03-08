package service

import (
	"context"

	"github.com/rdl/core/internal/modules/payment/domain"
)

type PaymentService interface {
	InitiatePayment(ctx context.Context, userID string, req *domain.CreatePaymentRequest) (*domain.Payment, error)
	GetByID(ctx context.Context, id string) (*domain.Payment, error)
	List(ctx context.Context, filter domain.ListPaymentsFilter) ([]*domain.Payment, error)
	HandleWebhook(ctx context.Context, req *domain.PaymentWebhookRequest) error
}
