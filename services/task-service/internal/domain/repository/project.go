package repository

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/google/uuid"
)

type ProjectRepository interface {
	Save(ctx context.Context, project *domain.Project) error
	List(ctx context.Context, limit, offset uint64) ([]*domain.Project, error)
	ListByOwner(ctx context.Context, ownerID uuid.UUID) ([]*domain.Project, error)
	GetById(ctx context.Context, projectID uuid.UUID) (*domain.Project, error)
	GetByOwner(ctx context.Context, ownerID uuid.UUID, projectID uuid.UUID) (*domain.Project, error)
	Delete(ctx context.Context, projectID uuid.UUID) error
	Update(ctx context.Context, project *domain.Project) error
}
