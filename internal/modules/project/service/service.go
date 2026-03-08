package service

import (
	"context"

	"github.com/rdl/core/internal/modules/project/domain"
)

type ProjectService interface {
	Create(ctx context.Context, ownerID string, req *domain.CreateProjectRequest) (*domain.Project, error)
	GetByID(ctx context.Context, id string) (*domain.Project, error)
	List(ctx context.Context, filter domain.ListProjectsFilter) ([]*domain.Project, error)
	Update(ctx context.Context, id string, req *domain.UpdateProjectRequest) (*domain.Project, error)
	Delete(ctx context.Context, id string) error
}
