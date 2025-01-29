package repository

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/google/uuid"
)

type TaskRepository interface {
	Save(ctx context.Context, task *domain.Task) error
	List(ctx context.Context, limit, offset uint64) ([]*domain.Task, error)
	ListByProject(ctx context.Context, projectID uuid.UUID) ([]*domain.Task, error)
	GetById(ctx context.Context, taskID uuid.UUID) (*domain.Task, error)
	Delete(ctx context.Context, taskID uuid.UUID) error
	Update(ctx context.Context, task *domain.Task) error
}
