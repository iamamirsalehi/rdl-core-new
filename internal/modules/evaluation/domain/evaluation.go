package domain

import "time"

type Evaluation struct {
	ID          string    `json:"id" bson:"_id"`
	ProjectID   string    `json:"project_id" bson:"project_id"`
	EvaluatorID string    `json:"evaluator_id" bson:"evaluator_id"`
	Score       float64   `json:"score" bson:"score"`
	Feedback    string    `json:"feedback" bson:"feedback"`
	Status      string    `json:"status" bson:"status"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type CreateEvaluationRequest struct {
	ProjectID string  `json:"project_id" validate:"required"`
	Score     float64 `json:"score"      validate:"required,min=0,max=10"`
	Feedback  string  `json:"feedback"`
}

type UpdateEvaluationRequest struct {
	Score    *float64 `json:"score"`
	Feedback *string  `json:"feedback"`
	Status   *string  `json:"status"`
}

type ListEvaluationsFilter struct {
	ProjectID   string
	EvaluatorID string
	Page        int
	Limit       int
}
