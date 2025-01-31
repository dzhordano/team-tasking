package services

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/application/interfaces"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain/repository"
	"github.com/google/uuid"
)

type projectService struct {
	// log *slog.Logger
	projectRepository repository.ProjectRepository
}

func NewProjectService(projectRepository repository.ProjectRepository) interfaces.ProjectService {
	return &projectService{
		projectRepository: projectRepository,
	}
}

func (s *projectService) CreateProject(ctx context.Context, title string, ownerID uuid.UUID) error {
	projectID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	project := domain.NewProject(projectID, ownerID, title)

	if err := project.Validate(); err != nil {
		return err
	}

	if err := s.projectRepository.Save(ctx, project); err != nil {
		return err
	}

	return nil
}

func (s *projectService) GetUserProjects(ctx context.Context, userID uuid.UUID) ([]*domain.Project, error) {
	projects, err := s.projectRepository.GetByOwner(ctx, userID)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (s *projectService) DeleteProject(ctx context.Context, ownerID, projectID uuid.UUID) error {
	if err := s.projectRepository.Delete(ctx, projectID); err != nil {
		return err
	}

	return nil
}
