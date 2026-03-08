package service

import (
	"context"

	"github.com/rdl/core/internal/modules/cathub/domain"
	"github.com/rdl/core/internal/modules/cathub/repository"
)

type cathubService struct {
	hubRepo repository.HubRepository
	catRepo repository.CategoryRepository
}

func New(hubRepo repository.HubRepository, catRepo repository.CategoryRepository) CathubService {
	return &cathubService{hubRepo: hubRepo, catRepo: catRepo}
}

func (s *cathubService) CreateHub(ctx context.Context, ownerID string, req *domain.CreateHubRequest) (*domain.Hub, error) {
	hub := &domain.Hub{
		CategoryID:  req.CategoryID,
		Title:       req.Title,
		Description: req.Description,
		OwnerID:     ownerID,
	}
	if err := s.hubRepo.Create(ctx, hub); err != nil {
		return nil, err
	}
	return hub, nil
}

func (s *cathubService) GetHubByID(ctx context.Context, id string) (*domain.Hub, error) {
	return s.hubRepo.FindByID(ctx, id)
}

func (s *cathubService) ListHubs(ctx context.Context, filter domain.ListHubsFilter) ([]*domain.Hub, error) {
	return s.hubRepo.List(ctx, filter)
}

func (s *cathubService) UpdateHub(ctx context.Context, id string, req *domain.UpdateHubRequest) (*domain.Hub, error) {
	hub, err := s.hubRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if req.Title != nil {
		hub.Title = *req.Title
	}
	if req.Description != nil {
		hub.Description = *req.Description
	}
	if err := s.hubRepo.Update(ctx, hub); err != nil {
		return nil, err
	}
	return hub, nil
}

func (s *cathubService) DeleteHub(ctx context.Context, id string) error {
	return s.hubRepo.Delete(ctx, id)
}
