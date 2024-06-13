package domain

import (
	"context"
	"github.com/Mx1q/ppo_services/domain"
	"github.com/google/uuid"
)

type IIngredientTypeRepository interface {
	Create(ctx context.Context, ingredientType *domain.IngredientType) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.IngredientType, error)
	GetAll(ctx context.Context) ([]*domain.IngredientType, error)
	Update(ctx context.Context, measurement *domain.IngredientType) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
