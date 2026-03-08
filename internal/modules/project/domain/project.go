package domain

import "time"

type Project struct {
	ID          string    `json:"id" bson:"_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	OwnerID     string    `json:"owner_id" bson:"owner_id"`
	Status      string    `json:"status" bson:"status"`
	Tags        []string  `json:"tags" bson:"tags"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type CreateProjectRequest struct {
	Title       string   `json:"title"       validate:"required"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type UpdateProjectRequest struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Status      *string  `json:"status"`
	Tags        []string `json:"tags"`
}

type ListProjectsFilter struct {
	OwnerID string
	Status  string
	Page    int
	Limit   int
}
