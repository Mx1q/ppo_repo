package domain

import (
	"context"
	"github.com/Mx1q/ppo_services/domain"
	"github.com/google/uuid"
)

type ISaladRepository interface {
	Create(ctx context.Context, salad *domain.Salad) (uuid.UUID, error)
	GetById(ctx context.Context, id uuid.UUID) (*domain.Salad, error)
	GetAll(ctx context.Context, filter *domain.RecipeFilter, page int) ([]*domain.Salad, int, error)
	GetAllByUserId(ctx context.Context, id uuid.UUID) ([]*domain.Salad, error)
	GetAllRatedByUser(ctx context.Context, userId uuid.UUID, page int) ([]*domain.Salad, int, error)
	Update(ctx context.Context, salad *domain.Salad) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
