package service

import (
	"context"

	"github.com/rdl/core/internal/modules/project/domain"
	"github.com/rdl/core/internal/modules/project/repository"
)

type projectService struct {
	repo repository.ProjectRepository
}

func New(repo repository.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) Create(ctx context.Context, ownerID string, req *domain.CreateProjectRequest) (*domain.Project, error) {
	p := &domain.Project{
		Title:       req.Title,
		Description: req.Description,
		OwnerID:     ownerID,
		Status:      "active",
		Tags:        req.Tags,
	}
	if err := s.repo.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *projectService) GetByID(ctx context.Context, id string) (*domain.Project, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *projectService) List(ctx context.Context, filter domain.ListProjectsFilter) ([]*domain.Project, error) {
	return s.repo.List(ctx, filter)
}

func (s *projectService) Update(ctx context.Context, id string, req *domain.UpdateProjectRequest) (*domain.Project, error) {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if req.Title != nil {
		p.Title = *req.Title
	}
	if req.Description != nil {
		p.Description = *req.Description
	}
	if req.Status != nil {
		p.Status = *req.Status
	}
	if req.Tags != nil {
		p.Tags = req.Tags
	}
	if err := s.repo.Update(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *projectService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
