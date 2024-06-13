package domain

import (
	"context"
	"github.com/Mx1q/ppo_services/domain"
	"github.com/google/uuid"
)

type IKeywordValidatorRepository interface {
	Create(ctx context.Context, word *domain.KeyWord) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.KeyWord, error)
	GetAll(ctx context.Context) (map[string]uuid.UUID, error)
	Update(ctx context.Context, word *domain.KeyWord) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
