package service

import (
	"context"
	"github.com/vasaud880/graden/pkg/domain"
	"github.com/vasaud880/graden/pkg/repository"
)

type PlanService struct {
	repo *repository.PlanRepository
}

func NewPlanService(repo *repository.PlanRepository) *PlanService {
	return &PlanService{repo: repo}
}

func (s *PlanService) CreatePlan(ctx context.Context, chatID int64) (*domain.CityPlan, error) {
	plan := &domain.CityPlan{}
	err := s.repo.Save(ctx, chatID, plan)
	return plan, err
}

func (s *PlanService) GetPlan(ctx context.Context, chatID int64) (*domain.CityPlan, error) {
	return s.repo.Get(ctx, chatID)
}

func (s *PlanService) UpdatePlanName(ctx context.Context, chatID int64, name string) error {
	plan, err := s.repo.Get(ctx, chatID)
	if err != nil {
		return err
	}
	plan.Name = name
	return s.repo.Save(ctx, chatID, plan)
}

func (s *PlanService) UpdatePlanDescription(ctx context.Context, chatID int64, description string) error {
	plan, err := s.repo.Get(ctx, chatID)
	if err != nil {
		return err
	}
	plan.Description = description
	return s.repo.Save(ctx, chatID, plan)
}

func (s *PlanService) UpdatePlanArea(ctx context.Context, chatID int64, area float64) error {
	plan, err := s.repo.Get(ctx, chatID)
	if err != nil {
		return err
	}
	plan.Area = area
	return s.repo.Save(ctx, chatID, plan)
}

func (s *PlanService) UpdatePlanPopulation(ctx context.Context, chatID int64, population int) error {
	plan, err := s.repo.Get(ctx, chatID)
	if err != nil {
		return err
	}
	plan.Population = population
	return s.repo.Save(ctx, chatID, plan)
}

func (s *PlanService) DeletePlan(ctx context.Context, chatID int64) error {
	return s.repo.Delete(ctx, chatID)
}
