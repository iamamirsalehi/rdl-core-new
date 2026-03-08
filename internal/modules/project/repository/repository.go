package repository

import (
	"context"

	"github.com/rdl/core/internal/modules/project/domain"
)

type ProjectRepository interface {
	Create(ctx context.Context, project *domain.Project) error
	FindByID(ctx context.Context, id string) (*domain.Project, error)
	List(ctx context.Context, filter domain.ListProjectsFilter) ([]*domain.Project, error)
	Update(ctx context.Context, project *domain.Project) error
	Delete(ctx context.Context, id string) error
}
