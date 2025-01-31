package repository

import (
	"context"
	"time"

	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/google/uuid"
)

type TaskRepository interface {
	Save(ctx context.Context, task *domain.Task) error
	List(ctx context.Context, limit, offset uint64) ([]*domain.Task, error)
	ListByProject(ctx context.Context, projectID uuid.UUID) ([]*domain.Task, error)
	GetById(ctx context.Context, taskID uuid.UUID) (*domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
	GetByUserId(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error)
	// Delete(ctx context.Context, taskID uuid.UUID) error

	AssignTask(ctx context.Context, taskID, assigneeID uuid.UUID, created_at time.Time) error
	ListPendingTasks(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*domain.Task, error)
	AcceptTask(ctx context.Context, taskID, userID uuid.UUID) error
	DeclineTask(ctx context.Context, taskID, userID uuid.UUID) error
}
