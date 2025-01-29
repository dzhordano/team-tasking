package repository

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGTaskRepository struct {
	db *pgxpool.Pool
}

func NewPGTaskRepository(db *pgxpool.Pool) repository.TaskRepository {
	return &PGTaskRepository{
		db: db,
	}
}

func (t *PGTaskRepository) Save(ctx context.Context, task *domain.Task) error {
	return nil
}

func (t *PGTaskRepository) List(ctx context.Context, limit, offset uint64) ([]*domain.Task, error) {
	return nil, nil
}

func (t *PGTaskRepository) ListByProject(ctx context.Context, projectID uuid.UUID) ([]*domain.Task, error) {
	return nil, nil
}

func (t *PGTaskRepository) GetById(ctx context.Context, taskID uuid.UUID) (*domain.Task, error) {
	return nil, nil
}

func (t *PGTaskRepository) Update(ctx context.Context, task *domain.Task) error {
	return nil
}

func (t *PGTaskRepository) Delete(ctx context.Context, taskID uuid.UUID) error {
	return nil
}
