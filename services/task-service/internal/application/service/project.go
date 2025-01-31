package services

import (
	"context"
	"log/slog"

	"github.com/dzhordano/team-tasking/services/tasks/internal/application/interfaces"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain/repository"
	"github.com/google/uuid"
)

type projectService struct {
	log               *slog.Logger
	projectRepository repository.ProjectRepository
}

func NewProjectService(log *slog.Logger, projectRepository repository.ProjectRepository) interfaces.ProjectService {
	return &projectService{
		log:               log,
		projectRepository: projectRepository,
	}
}

func (s *projectService) CreateProject(ctx context.Context, title string, ownerID uuid.UUID) error {
	projectID, err := uuid.NewUUID()
	if err != nil {
		s.log.Error("failed to generate project id", slog.String("error", err.Error()))
		return err
	}

	project := domain.NewProject(projectID, ownerID, title)

	if err := project.Validate(); err != nil {
		s.log.Error("failed to validate project", slog.String("error", err.Error()))
		return err
	}

	if err := s.projectRepository.Save(ctx, project); err != nil {
		s.log.Error("failed to save project", slog.String("error", err.Error()))
		return err
	}

	s.log.Debug("project created", slog.String("project_id", projectID.String()))

	return nil
}

func (s *projectService) GetUserProjects(ctx context.Context, userID uuid.UUID) ([]*domain.Project, error) {
	projects, err := s.projectRepository.ListByOwner(ctx, userID)
	if err != nil {
		s.log.Error("failed to get user projects", slog.String("error", err.Error()))
		return nil, err
	}

	if len(projects) == 0 {
		s.log.Debug("no projects found for user", slog.String("user_id", userID.String()))
		return nil, domain.ErrNoProjectsFound
	}

	s.log.Debug("user projects found", slog.String("user_id", userID.String()))

	return projects, nil
}

func (s *projectService) DeleteProject(ctx context.Context, userID, projectID uuid.UUID) error {
	_, err := s.projectRepository.GetByOwner(ctx, userID, projectID)
	if err != nil {
		s.log.Error("project not found", slog.String("project_id", projectID.String()))
		return err
	}

	if err := s.projectRepository.Delete(ctx, projectID); err != nil {
		s.log.Error("failed to delete project", slog.String("error", err.Error()))
		return err
	}

	s.log.Debug("project deleted", slog.String("project_id", projectID.String()))

	return nil
}
