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
	tasks_table = "tasks"
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
	const op = "repository.PGTaskRepository.Save"

	insertBuild := sq.Insert(tasks_table).
		Columns("id", "project_id", "assignee_id", "title", "description", "status", "deadline", "created_at", "updated_at").
		Values(task.TaskID, task.ProjectID, task.AssigneeID, task.Title, task.Description, task.Status, task.Deadline, task.CreatedAt, task.UpdatedAt).
		PlaceholderFormat(sq.Dollar)

	query, args, err := insertBuild.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := t.db.Exec(ctx, query, args...); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return fmt.Errorf("%s: %w", op, domain.ErrTaskAlreadyExists)
			}
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (t *PGTaskRepository) List(ctx context.Context, limit, offset uint64) ([]*domain.Task, error) {
	return nil, nil
}

func (t *PGTaskRepository) ListByProject(ctx context.Context, projectID uuid.UUID) ([]*domain.Task, error) {
	return nil, nil
}

func (t *PGTaskRepository) GetById(ctx context.Context, taskID uuid.UUID) (*domain.Task, error) {
	const op = "repository.PGTaskRepository.GetById"

	selectBuilder := sq.Select("id", "project_id", "assignee_id", "title", "description", "status", "deadline", "created_at", "updated_at").
		From(tasks_table).
		Where(sq.Eq{"id": taskID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var task domain.Task
	if err := t.db.QueryRow(ctx, query, args...).Scan(&task.TaskID, &task.ProjectID, &task.AssigneeID, &task.Title, &task.Description, &task.Status, &task.Deadline, &task.CreatedAt, &task.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, domain.ErrTaskNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &task, nil
}

func (t *PGTaskRepository) GetByUserId(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error) {
	return nil, nil
}

func (t *PGTaskRepository) Update(ctx context.Context, task *domain.Task) error {
	return nil
}

func (t *PGTaskRepository) Delete(ctx context.Context, taskID uuid.UUID) error {
	return nil
}
