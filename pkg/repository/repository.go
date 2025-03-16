package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vasaud880/graden/pkg/domain"
)

var (
	ErrPlanNotFound = errors.New("plan not found")
)

type PlanRepository struct {
	db *pgxpool.Pool
}

func NewPlanRepository(db *pgxpool.Pool) *PlanRepository {
	return &PlanRepository{db: db}
}

func (r *PlanRepository) Save(ctx context.Context, chatID int64, plan *domain.CityPlan) error {
	query := `
		INSERT INTO city_plans (chat_id, name, description, area, population)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (chat_id) DO UPDATE
		SET name = $2, description = $3, area = $4, population = $5
	`
	_, err := r.db.Exec(ctx, query, chatID, plan.Name, plan.Description, plan.Area, plan.Population)
	return err
}

func (r *PlanRepository) Get(ctx context.Context, chatID int64) (*domain.CityPlan, error) {
	query := `
		SELECT name, description, area, population
		FROM city_plans
		WHERE chat_id = $1
	`
	var plan domain.CityPlan
	err := r.db.QueryRow(ctx, query, chatID).Scan(&plan.Name, &plan.Description, &plan.Area, &plan.Population)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrPlanNotFound
		}
		return nil, err
	}
	return &plan, nil
}

func (r *PlanRepository) Delete(ctx context.Context, chatID int64) error {
	query := `DELETE FROM city_plans WHERE chat_id = $1`
	_, err := r.db.Exec(ctx, query, chatID)
	return err
}
