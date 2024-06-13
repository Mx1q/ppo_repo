package domain

import (
	"context"
	"github.com/Mx1q/ppo_services/domain"
	"github.com/google/uuid"
)

type ICommentRepository interface {
	Create(ctx context.Context, comment *domain.Comment) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.Comment, error)
	GetBySaladAndUser(ctx context.Context, saladId uuid.UUID, userId uuid.UUID) (*domain.Comment, error)
	GetAllBySaladID(ctx context.Context, saladId uuid.UUID, page int) ([]*domain.Comment, int, error)
	Update(ctx context.Context, comment *domain.Comment) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
