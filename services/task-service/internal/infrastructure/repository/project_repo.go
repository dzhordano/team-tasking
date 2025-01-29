package repository

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGProjectRepository struct {
	db *pgxpool.Pool
}

func NewPGProjectRepository(db *pgxpool.Pool) repository.ProjectRepository {
	return &PGProjectRepository{
		db: db,
	}
}

func (p *PGProjectRepository) Save(ctx context.Context, project *domain.Project) error {
	return nil
}

func (p *PGProjectRepository) List(ctx context.Context, limit, offset uint64) ([]*domain.Project, error) {
	return nil, nil
}

func (p *PGProjectRepository) GetById(ctx context.Context, projectID uuid.UUID) (*domain.Project, error) {
	return nil, nil
}

func (p *PGProjectRepository) GetByOwner(ctx context.Context, ownerID uuid.UUID) ([]*domain.Project, error) {
	return nil, nil
}

func (p *PGProjectRepository) Update(ctx context.Context, project *domain.Project) error {
	return nil
}

func (p *PGProjectRepository) Delete(ctx context.Context, projectID uuid.UUID) error {
	return nil
}
