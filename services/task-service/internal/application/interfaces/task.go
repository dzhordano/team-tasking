package interfaces

import (
	"context"
	"time"

	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(ctx context.Context, title, description string, userId, projectID uuid.UUID, deadline time.Time) error
	AssignTask(ctx context.Context, taskID, assigneeID uuid.UUID) error
	GetUserTasks(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error)
}
