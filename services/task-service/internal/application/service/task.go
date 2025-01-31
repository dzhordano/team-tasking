package services

import (
	"context"
	"log/slog"
	"time"

	"github.com/dzhordano/team-tasking/services/tasks/internal/application/interfaces"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain/repository"
	"github.com/google/uuid"
)

type taskService struct {
	log               *slog.Logger
	taskRepository    repository.TaskRepository
	projectRepository repository.ProjectRepository
}

func NewTaskService(log *slog.Logger, taskRepository repository.TaskRepository, projectRepository repository.ProjectRepository) interfaces.TaskService {
	return &taskService{
		log:               log,
		taskRepository:    taskRepository,
		projectRepository: projectRepository,
	}
}

func (s *taskService) CreateTask(ctx context.Context, title, description string, userId, projectID uuid.UUID, deadline time.Time) error {
	_, err := s.projectRepository.GetByOwner(ctx, userId, projectID)
	if err != nil {
		s.log.Error("project not found", slog.String("project_id", projectID.String()))
		return err
	}

	taskID, err := uuid.NewUUID()
	if err != nil {
		s.log.Error("failed to generate task id", slog.String("error", err.Error()))
		return err
	}

	task := domain.NewTask(taskID, projectID, title, description, deadline)

	if err := task.Validate(); err != nil {
		s.log.Error("failed to validate task", slog.String("error", err.Error()))
		return err
	}

	if err := s.taskRepository.Save(ctx, task); err != nil {
		s.log.Error("failed to save task", slog.String("error", err.Error()))
		return err
	}

	s.log.Debug("task created", slog.String("task_id", taskID.String()))

	return nil
}

func (s *taskService) AssignTask(ctx context.Context, userID, taskID, assigneeID uuid.UUID) error {
	task, err := s.taskRepository.GetById(ctx, taskID)
	if err != nil {
		s.log.Error("task not found", slog.String("task_id", taskID.String()))
		return err
	}

	_, err = s.projectRepository.GetByOwner(ctx, userID, task.ProjectID)
	if err != nil {
		s.log.Error("project not found", slog.String("project_id", taskID.String()))
		return err
	}

	task.SetAssignee(assigneeID)

	if err := s.taskRepository.AssignTask(ctx, taskID, assigneeID, time.Now()); err != nil {
		s.log.Error("failed to assign task", slog.String("error", err.Error()))
		return err
	}

	s.log.Debug("task assigned", slog.String("task_id", taskID.String()))

	return nil
}

func (s *taskService) GetUserTasks(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error) {
	tasks, err := s.taskRepository.GetByUserId(ctx, userID)
	if err != nil {
		s.log.Error("failed to get user tasks", slog.String("error", err.Error()))
		return nil, err
	}

	if len(tasks) == 0 {
		s.log.Debug("no tasks found for user", slog.String("user_id", userID.String()))
		return nil, domain.ErrNoTasksFound
	}

	s.log.Debug("user tasks found", slog.String("user_id", userID.String()))

	return tasks, nil
}

func (s *taskService) GetPendingTasks(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*domain.Task, error) {
	if limit == 0 {
		limit = 10 // FIXME magic number
	}

	tasks, err := s.taskRepository.ListPendingTasks(ctx, userID, limit, offset)
	if err != nil {
		s.log.Error("failed to get pending tasks", slog.String("error", err.Error()))
		return nil, err
	}

	if len(tasks) == 0 {
		s.log.Debug("no pending tasks found for user", slog.String("user_id", userID.String()))
		return nil, domain.ErrNoPendingTasksFound
	}

	s.log.Debug("pending tasks found", slog.String("user_id", userID.String()))

	return tasks, nil
}

func (s *taskService) AcceptTask(ctx context.Context, taskID, userID uuid.UUID) error {
	task, err := s.taskRepository.GetById(ctx, taskID)
	if err != nil {
		s.log.Error("task not found", slog.String("task_id", taskID.String()))
		return err
	}

	if err := s.taskRepository.AcceptTask(ctx, taskID, userID); err != nil {
		s.log.Error("failed to accept task", slog.String("error", err.Error()))
		return err
	}

	task.SetAssignee(userID)
	task.SetStatus(domain.TaskINPROGRESS)

	if err := s.taskRepository.Update(ctx, task); err != nil {
		s.log.Error("failed to update task", slog.String("error", err.Error()))
		return err
	}

	s.log.Debug("task accepted", slog.String("task_id", taskID.String()))

	return nil
}

func (s *taskService) DeclineTask(ctx context.Context, taskID, userID uuid.UUID) error {
	_, err := s.taskRepository.GetById(ctx, taskID)
	if err != nil {
		s.log.Error("task not found", slog.String("task_id", taskID.String()))
		return err
	}

	if err := s.taskRepository.DeclineTask(ctx, taskID, userID); err != nil {
		s.log.Error("failed to decline task", slog.String("error", err.Error()))
		return err
	}

	s.log.Debug("task declined", slog.String("task_id", taskID.String()))

	return nil
}

func (s *taskService) FinishTask(ctx context.Context, taskID, userID uuid.UUID) error {
	task, err := s.taskRepository.GetById(ctx, taskID)
	if err != nil {
		s.log.Error("task not found", slog.String("task_id", taskID.String()))
		return err
	}

	if task.AssigneeID != userID {
		s.log.Error("permission denied", slog.String("task_id", taskID.String()))
		return domain.ErrTaskNotFound
	}

	task.SetStatus(domain.TaskDONE)

	if err := s.taskRepository.Update(ctx, task); err != nil {
		s.log.Error("failed to update task", slog.String("error", err.Error()))
		return err
	}

	s.log.Debug("task finished", slog.String("task_id", taskID.String()))

	return nil
}
