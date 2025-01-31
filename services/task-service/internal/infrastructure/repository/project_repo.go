package repository

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	projects_table = "projects"
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
	const op = "repository.PGProjectRepository.Save"

	insertBuilder := sq.Insert(projects_table).
		Columns("id", "owner_id", "name", "created_at", "updated_at").
		Values(project.ProjectID, project.OwnerID, project.Name, project.CreatedAt, project.UpdatedAt).
		PlaceholderFormat(sq.Dollar)

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	if _, err := p.db.Exec(ctx, query, args...); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return fmt.Errorf("%s : %w", op, domain.ErrProjectAlreadyExists)
			}
		}
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}

func (p *PGProjectRepository) List(ctx context.Context, limit, offset uint64) ([]*domain.Project, error) {
	return nil, nil
}

func (p *PGProjectRepository) ListByOwner(ctx context.Context, ownerID uuid.UUID) ([]*domain.Project, error) {
	return nil, nil
}

func (p *PGProjectRepository) GetById(ctx context.Context, projectID uuid.UUID) (*domain.Project, error) {
	return nil, nil
}

func (p *PGProjectRepository) GetByOwner(ctx context.Context, ownerID uuid.UUID, projectID uuid.UUID) (*domain.Project, error) {
	const op = "repository.PGProjectRepository.GetByOwner"

	selectBuilder := sq.Select("id", "owner_id", "name", "created_at", "updated_at").
		From(projects_table).
		Where(sq.Eq{"owner_id": ownerID, "id": projectID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	var project domain.Project
	if err := p.db.QueryRow(ctx, query, args...).Scan(&project.ProjectID, &project.OwnerID, &project.Name, &project.CreatedAt, &project.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s : %w", op, domain.ErrProjectNotFound)
		}
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	return &project, nil
}

func (p *PGProjectRepository) Update(ctx context.Context, project *domain.Project) error {
	return nil
}

func (p *PGProjectRepository) Delete(ctx context.Context, projectID uuid.UUID) error {
	return nil
}
