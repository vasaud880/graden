package repository

import (
	"errors"
	"sync"

	"github.com/vasaud880/graden/internal/domain"
)

var (
	ErrPlanNotFound = errors.New("plan not found")
)

type PlanRepository struct {
	plans map[int64]*domain.CityPlan
	mu    sync.RWMutex
}

func NewPlanRepository() *PlanRepository {
	return &PlanRepository{
		plans: make(map[int64]*domain.CityPlan),
	}
}

func (r *PlanRepository) Save(chatID int64, plan *domain.CityPlan) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.plans[chatID] = plan
}

func (r *PlanRepository) Get(chatID int64) (*domain.CityPlan, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	plan, exists := r.plans[chatID]
	if !exists {
		return nil, ErrPlanNotFound
	}
	return plan, nil
}

func (r *PlanRepository) Delete(chatID int64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.plans, chatID)
}
