package domain

import (
	"context"
	"github.com/Mx1q/ppo_services/domain"
	"github.com/google/uuid"
)

type IIngredientRepository interface {
	Create(ctx context.Context, ingredient *domain.Ingredient) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.Ingredient, error)
	GetAll(ctx context.Context, page int) ([]*domain.Ingredient, int, error)
	GetAllByRecipeId(ctx context.Context, id uuid.UUID) ([]*domain.Ingredient, error)
	Link(ctx context.Context, recipeId uuid.UUID, ingredientId uuid.UUID) (uuid.UUID, error)
	Unlink(ctx context.Context, recipeId uuid.UUID, ingredientId uuid.UUID) error
	Update(ctx context.Context, ingredient *domain.Ingredient) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
