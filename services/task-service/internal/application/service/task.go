package services

import (
	"context"
	"time"

	"github.com/dzhordano/team-tasking/services/tasks/internal/application/interfaces"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain/repository"
	"github.com/google/uuid"
)

type taskService struct {
	// log *slog.Logger
	taskRepository    repository.TaskRepository
	projectRepository repository.ProjectRepository
}

func NewTaskService(taskRepository repository.TaskRepository, projectRepository repository.ProjectRepository) interfaces.TaskService {
	return &taskService{
		taskRepository:    taskRepository,
		projectRepository: projectRepository,
	}
}

func (s *taskService) CreateTask(ctx context.Context, title, description string, userId, projectID uuid.UUID, deadline time.Time) error {
	_, err := s.projectRepository.GetByOwner(ctx, userId)
	if err != nil {
		return err
	}

	taskID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	task := domain.NewTask(taskID, projectID, title, description, deadline)

	if err := task.Validate(); err != nil {
		return err
	}

	if err := s.taskRepository.Save(ctx, task); err != nil {
		return err
	}

	return nil
}

func (s *taskService) AssignTask(ctx context.Context, taskID, assigneeID uuid.UUID) error {
	task, err := s.taskRepository.GetById(ctx, taskID)
	if err != nil {
		return err
	}

	task.SetAssignee(assigneeID)

	if err := s.taskRepository.Update(ctx, task); err != nil {
		return err
	}

	return nil
}

func (s *taskService) GetUserTasks(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error) {
	tasks, err := s.taskRepository.GetByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
