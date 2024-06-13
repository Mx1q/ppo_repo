package domain

import (
	"context"
	"github.com/Mx1q/ppo_services/domain"
	"github.com/google/uuid"
)

type IMeasurementRepository interface {
	Create(ctx context.Context, measurement *domain.Measurement) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.Measurement, error)
	GetByRecipeId(ctx context.Context, ingredientId uuid.UUID, recipeId uuid.UUID) (*domain.Measurement, int, error)
	GetAll(ctx context.Context) ([]*domain.Measurement, error)
	UpdateLink(ctx context.Context, linkId uuid.UUID, measurementId uuid.UUID, amount int) error
	Update(ctx context.Context, measurement *domain.Measurement) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
