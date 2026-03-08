package domain

import "time"

type Payment struct {
	ID         string    `json:"id" bson:"_id"`
	UserID     string    `json:"user_id" bson:"user_id"`
	Amount     float64   `json:"amount" bson:"amount"`
	Currency   string    `json:"currency" bson:"currency"`
	Status     string    `json:"status" bson:"status"`
	Provider   string    `json:"provider" bson:"provider"`
	ProviderID string    `json:"provider_id" bson:"provider_id"`
	Metadata   map[string]any `json:"metadata,omitempty" bson:"metadata,omitempty"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

type CreatePaymentRequest struct {
	Amount   float64 `json:"amount"   validate:"required,gt=0"`
	Currency string  `json:"currency" validate:"required,len=3"`
	Provider string  `json:"provider" validate:"required"`
}

type PaymentWebhookRequest struct {
	Provider string         `json:"provider"`
	Payload  map[string]any `json:"payload"`
}

type ListPaymentsFilter struct {
	UserID   string
	Status   string
	Provider string
	Page     int
	Limit    int
}
