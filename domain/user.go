package domain

import (
	"context"
	"github.com/Mx1q/ppo_services/domain"
	"github.com/google/uuid"
)

type IUserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	GetAll(ctx context.Context, page int) ([]*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
