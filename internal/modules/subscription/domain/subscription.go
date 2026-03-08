package domain

import "time"

type Plan struct {
	ID          string    `json:"id" bson:"_id"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Price       float64   `json:"price" bson:"price"`
	Currency    string    `json:"currency" bson:"currency"`
	Interval    string    `json:"interval" bson:"interval"` // monthly, yearly
	Features    []string  `json:"features" bson:"features"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}

type Subscription struct {
	ID         string    `json:"id" bson:"_id"`
	UserID     string    `json:"user_id" bson:"user_id"`
	PlanID     string    `json:"plan_id" bson:"plan_id"`
	Status     string    `json:"status" bson:"status"` // active, cancelled, expired
	StartedAt  time.Time `json:"started_at" bson:"started_at"`
	ExpiresAt  time.Time `json:"expires_at" bson:"expires_at"`
	CancelledAt *time.Time `json:"cancelled_at,omitempty" bson:"cancelled_at,omitempty"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

type CreateSubscriptionRequest struct {
	PlanID string `json:"plan_id" validate:"required"`
}

type ListSubscriptionsFilter struct {
	UserID string
	Status string
	Page   int
	Limit  int
}
