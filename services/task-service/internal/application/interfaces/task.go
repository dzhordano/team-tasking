package interfaces

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(ctx context.Context, title, description string, projectID uuid.UUID) error
	AssignTask(ctx context.Context, taskID, assigneeID uuid.UUID) error
	GetUserTasks(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error)
}
