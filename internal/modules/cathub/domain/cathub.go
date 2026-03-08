package domain

import "time"

type Category struct {
	ID          string    `json:"id" bson:"_id"`
	Name        string    `json:"name" bson:"name"`
	Slug        string    `json:"slug" bson:"slug"`
	Description string    `json:"description" bson:"description"`
	ParentID    string    `json:"parent_id,omitempty" bson:"parent_id,omitempty"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type Hub struct {
	ID          string    `json:"id" bson:"_id"`
	CategoryID  string    `json:"category_id" bson:"category_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	OwnerID     string    `json:"owner_id" bson:"owner_id"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type CreateHubRequest struct {
	CategoryID  string `json:"category_id"  validate:"required"`
	Title       string `json:"title"        validate:"required"`
	Description string `json:"description"`
}

type UpdateHubRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type ListHubsFilter struct {
	CategoryID string
	OwnerID    string
	Page       int
	Limit      int
}
