package service

import (
	"context"

	"github.com/rdl/core/internal/modules/evaluation/domain"
	"github.com/rdl/core/internal/modules/evaluation/repository"
)

type evaluationService struct {
	repo repository.EvaluationRepository
}

func New(repo repository.EvaluationRepository) EvaluationService {
	return &evaluationService{repo: repo}
}

func (s *evaluationService) Create(ctx context.Context, evaluatorID string, req *domain.CreateEvaluationRequest) (*domain.Evaluation, error) {
	e := &domain.Evaluation{
		ProjectID:   req.ProjectID,
		EvaluatorID: evaluatorID,
		Score:       req.Score,
		Feedback:    req.Feedback,
		Status:      "pending",
	}
	if err := s.repo.Create(ctx, e); err != nil {
		return nil, err
	}
	return e, nil
}

func (s *evaluationService) GetByID(ctx context.Context, id string) (*domain.Evaluation, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *evaluationService) List(ctx context.Context, filter domain.ListEvaluationsFilter) ([]*domain.Evaluation, error) {
	return s.repo.List(ctx, filter)
}

func (s *evaluationService) Update(ctx context.Context, id string, req *domain.UpdateEvaluationRequest) (*domain.Evaluation, error) {
	e, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if req.Score != nil {
		e.Score = *req.Score
	}
	if req.Feedback != nil {
		e.Feedback = *req.Feedback
	}
	if req.Status != nil {
		e.Status = *req.Status
	}
	if err := s.repo.Update(ctx, e); err != nil {
		return nil, err
	}
	return e, nil
}

func (s *evaluationService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
