package services

import (
	"context"
	"fmt"

	"github.com/dzhordano/team-tasking/services/tasks/internal/application/interfaces"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain/repository"
	"github.com/google/uuid"
)

type projectService struct {
	projectRepository repository.ProjectRepository
}

func NewProjectService(projectRepository repository.ProjectRepository) interfaces.ProjectService {
	return &projectService{
		projectRepository: projectRepository,
	}
}

func (s *projectService) CreateProject(ctx context.Context, title string, ownerID uuid.UUID) error {
	fmt.Println("You called CreateProject method!")

	return nil
}

func (s *projectService) ListProjects(ctx context.Context) ([]*domain.Project, error) {
	panic("TODO")
}

func (s *projectService) GetUserProjects(ctx context.Context, userID uuid.UUID) ([]*domain.Project, error) {
	panic("TODO")
}

func (s *projectService) DeleteProject(ctx context.Context, ownerID, projectID uuid.UUID) error {
	panic("TODO")
}
