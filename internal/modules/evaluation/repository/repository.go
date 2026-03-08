package repository

import (
	"context"

	"github.com/rdl/core/internal/modules/evaluation/domain"
)

type EvaluationRepository interface {
	Create(ctx context.Context, evaluation *domain.Evaluation) error
	FindByID(ctx context.Context, id string) (*domain.Evaluation, error)
	List(ctx context.Context, filter domain.ListEvaluationsFilter) ([]*domain.Evaluation, error)
	Update(ctx context.Context, evaluation *domain.Evaluation) error
	Delete(ctx context.Context, id string) error
}
