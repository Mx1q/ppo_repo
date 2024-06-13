package domain

import (
	"context"
	"github.com/Mx1q/ppo_services/domain"
	"github.com/google/uuid"
)

type IRecipeRepository interface {
	Create(ctx context.Context, recipe *domain.Recipe) (uuid.UUID, error)
	GetById(ctx context.Context, id uuid.UUID) (*domain.Recipe, error)
	GetBySaladId(ctx context.Context, saladId uuid.UUID) (*domain.Recipe, error)
	GetAll(ctx context.Context, filter *domain.RecipeFilter, page int) ([]*domain.Recipe, error)
	Update(ctx context.Context, recipe *domain.Recipe) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
