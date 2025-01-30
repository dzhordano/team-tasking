package interfaces

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/google/uuid"
)

type ProjectService interface {
	CreateProject(ctx context.Context, name string, ownerID uuid.UUID) error
	GetUserProjects(ctx context.Context, userID uuid.UUID) ([]*domain.Project, error)
	DeleteProject(ctx context.Context, ownerID, projectID uuid.UUID) error
}
