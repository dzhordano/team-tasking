package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	tasks_table         = "tasks"
	pending_tasks_table = "pending_tasks"
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

func (t *PGTaskRepository) Update(ctx context.Context, task *domain.Task) error {
	const op = "repository.PGTaskRepository.Update"

	updateBuilder := sq.Update(tasks_table).
		Where(sq.Eq{"id": task.TaskID}).
		Set("project_id", task.ProjectID).
		Set("assignee_id", task.AssigneeID).
		Set("title", task.Title).
		Set("description", task.Description).
		Set("status", task.Status).
		Set("deadline", task.Deadline).
		Set("updated_at", task.UpdatedAt).
		PlaceholderFormat(sq.Dollar)

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := t.db.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (t *PGTaskRepository) GetByUserId(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error) {
	const op = "repository.PGTaskRepository.GetByUserId"

	selectBuilder := sq.Select("id", "project_id", "assignee_id", "title", "description", "status", "deadline", "created_at", "updated_at").
		From(tasks_table).
		Where(sq.Eq{"assignee_id": userID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := t.db.Query(ctx, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, domain.ErrNoTasksFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var tasks []*domain.Task
	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.TaskID, &task.ProjectID, &task.AssigneeID, &task.Title, &task.Description, &task.Status, &task.Deadline, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

// func (t *PGTaskRepository) Delete(ctx context.Context, taskID uuid.UUID) error {
// 	return nil
// }

func (t *PGTaskRepository) AssignTask(ctx context.Context, taskID, assigneeID uuid.UUID, created_at time.Time) error {
	const op = "repository.PGTaskRepository.AssignTask"

	insertQuery := sq.Insert(pending_tasks_table).
		Columns("task_id", "user_id", "created_at").
		Values(taskID, assigneeID, created_at).
		PlaceholderFormat(sq.Dollar)

	query, args, err := insertQuery.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := t.db.Exec(ctx, query, args...); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return fmt.Errorf("%s: %w", op, domain.ErrTaskRequestAlreadyExists)
			}
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (t *PGTaskRepository) ListPendingTasks(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*domain.Task, error) {
	const op = "repository.PGTaskRepository.ListPendingTasks"

	fmt.Println("HERE")

	selectQuery := sq.Select("t.id", "t.project_id", "t.assignee_id", "t.title", "t.description", "t.status", "t.deadline", "t.created_at", "t.updated_at").
		From(pending_tasks_table).
		Join(tasks_table + " t on t.id = pending_tasks.task_id").
		Where(sq.Eq{"assignee_id": userID}).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectQuery.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := t.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var tasks []*domain.Task
	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.TaskID, &task.ProjectID, &task.AssigneeID, &task.Title, &task.Description, &task.Status, &task.Deadline, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t *PGTaskRepository) AcceptTask(ctx context.Context, taskID, userID uuid.UUID) error {
	const op = "repository.PGTaskRepository.AcceptTask"

	deleteQuery := sq.Delete(pending_tasks_table).
		Where(sq.Eq{"task_id": taskID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := deleteQuery.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := t.db.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (t *PGTaskRepository) DeclineTask(ctx context.Context, taskID, userID uuid.UUID) error {
	const op = "repository.PGTaskRepository.DeclineTask"

	deleteQuery := sq.Delete(pending_tasks_table).
		Where(sq.Eq{"task_id": taskID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := deleteQuery.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := t.db.Exec(ctx, query, args...); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return fmt.Errorf("%s: %w", op, domain.ErrTaskRequestAlreadyExists)
			}
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
