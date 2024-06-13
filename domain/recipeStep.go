package domain

import (
	"context"
	"github.com/Mx1q/ppo_services/domain"
	"github.com/google/uuid"
)

type IRecipeStepRepository interface {
	Create(ctx context.Context, recipeStep *domain.RecipeStep) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.RecipeStep, error)
	GetAllByRecipeID(ctx context.Context, recipeId uuid.UUID) ([]*domain.RecipeStep, error)
	Update(ctx context.Context, recipeStep *domain.RecipeStep) error
	DeleteById(ctx context.Context, id uuid.UUID) error
	DeleteAllByRecipeID(ctx context.Context, recipeId uuid.UUID) error
}
