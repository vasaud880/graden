package service

import (
	"github.com/vasaud880/graden/pkg/domain"
	"github.com/vasaud880/graden/pkg/repository"
)

type PlanService struct {
	repo *repository.PlanRepository
}

func NewPlanService(repo *repository.PlanRepository) *PlanService {
	return &PlanService{repo: repo}
}

func (s *PlanService) CreatePlan(chatID int64) *domain.CityPlan {
	plan := &domain.CityPlan{}
	s.repo.Save(chatID, plan)
	return plan
}

func (s *PlanService) GetPlan(chatID int64) (*domain.CityPlan, error) {
	return s.repo.Get(chatID)
}

func (s *PlanService) UpdatePlanName(chatID int64, name string) error {
	plan, err := s.repo.Get(chatID)
	if err != nil {
		return err
	}
	plan.Name = name
	s.repo.Save(chatID, plan)
	return nil
}

func (s *PlanService) UpdatePlanDescription(chatID int64, description string) error {
	plan, err := s.repo.Get(chatID)
	if err != nil {
		return err
	}
	plan.Description = description
	s.repo.Save(chatID, plan)
	return nil
}

func (s *PlanService) UpdatePlanArea(chatID int64, area float64) error {
	plan, err := s.repo.Get(chatID)
	if err != nil {
		return err
	}
	plan.Area = area
	s.repo.Save(chatID, plan)
	return nil
}

func (s *PlanService) UpdatePlanPopulation(chatID int64, population int) error {
	plan, err := s.repo.Get(chatID)
	if err != nil {
		return err
	}
	plan.Population = population
	s.repo.Save(chatID, plan)
	return nil
}
