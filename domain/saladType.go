package domain

import (
	"context"
	"github.com/Mx1q/ppo_services/domain"
	"github.com/google/uuid"
)

type ISaladTypeRepository interface {
	Create(ctx context.Context, saladType *domain.SaladType) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.SaladType, error)
	GetAll(ctx context.Context, page int) ([]*domain.SaladType, int, error)
	GetAllBySaladId(ctx context.Context, saladId uuid.UUID) ([]*domain.SaladType, error)
	Update(ctx context.Context, saladType *domain.SaladType) error
	Link(ctx context.Context, saladId uuid.UUID, saladTypeId uuid.UUID) error
	Unlink(ctx context.Context, saladId uuid.UUID, saladTypeId uuid.UUID) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
