package services

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/application/interfaces"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain/repository"
	"github.com/google/uuid"
)

type taskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) interfaces.TaskService {
	return &taskService{
		taskRepository: taskRepository,
	}
}

func (s *taskService) CreateTask(ctx context.Context, title, description string, projectID uuid.UUID) error {
	panic("TODO")
}

func (s *taskService) AssignTask(ctx context.Context, taskID, assigneeID uuid.UUID) error {
	panic("TODO")
}

func (s *taskService) GetUserTasks(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error) {
	panic("TODO")
}
