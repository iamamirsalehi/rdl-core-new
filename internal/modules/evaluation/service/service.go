package service

import (
	"context"

	"github.com/rdl/core/internal/modules/evaluation/domain"
)

type EvaluationService interface {
	Create(ctx context.Context, evaluatorID string, req *domain.CreateEvaluationRequest) (*domain.Evaluation, error)
	GetByID(ctx context.Context, id string) (*domain.Evaluation, error)
	List(ctx context.Context, filter domain.ListEvaluationsFilter) ([]*domain.Evaluation, error)
	Update(ctx context.Context, id string, req *domain.UpdateEvaluationRequest) (*domain.Evaluation, error)
	Delete(ctx context.Context, id string) error
}
