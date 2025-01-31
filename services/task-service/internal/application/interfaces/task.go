package interfaces

import (
	"context"
	"time"

	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(ctx context.Context, title, description string, userId, projectID uuid.UUID, deadline time.Time) error
	AssignTask(ctx context.Context, userID, taskID, assigneeID uuid.UUID) error
	GetUserTasks(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error)

	GetPendingTasks(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*domain.Task, error)
	AcceptTask(ctx context.Context, taskID, userID uuid.UUID) error
	DeclineTask(ctx context.Context, taskID, userID uuid.UUID) error
	FinishTask(ctx context.Context, taskID, userID uuid.UUID) error
}
