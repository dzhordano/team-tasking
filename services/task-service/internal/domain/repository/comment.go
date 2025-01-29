package repository

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/google/uuid"
)

type CommentRepository interface {
	Save(ctx context.Context, comment *domain.Comment) error
	List(ctx context.Context, limit, offset uint64) ([]*domain.Comment, error)
	GetById(ctx context.Context, commentID uuid.UUID) (*domain.Comment, error)
	GetByTask(ctx context.Context, taskID uuid.UUID) ([]*domain.Comment, error)
	Delete(ctx context.Context, commentID uuid.UUID) error
	Update(ctx context.Context, comment *domain.Comment) error
}
