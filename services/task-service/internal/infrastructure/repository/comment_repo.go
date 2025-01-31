package repository

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	comments_table = "comments"
)

type PGCommentRepository struct {
	db *pgxpool.Pool
}

func NewPGCommentRepository(db *pgxpool.Pool) repository.CommentRepository {
	return &PGCommentRepository{
		db: db,
	}
}

func (r *PGCommentRepository) Save(ctx context.Context, comment *domain.Comment) error {
	const op = "repository.PGCommentRepository.Save"

	insertBuilder := sq.Insert(comments_table).
		Columns("id", "task_id", "author_id", "content", "created_at", "updated_at").
		Values(comment.CommentID, comment.TaskID, comment.AuthorID, comment.Content, comment.CreatedAt, comment.UpdatedAt).
		PlaceholderFormat(sq.Dollar)

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}

func (r *PGCommentRepository) GetUserComments(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*domain.Comment, error) {
	const op = "repository.PGCommentRepository.GetUserComments"

	selectBuilder := sq.Select("id", "task_id", "author_id", "content", "created_at", "updated_at").
		From(comments_table).
		Where(sq.Eq{"author_id": userID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNoCommentsFound
		}
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	var comments []*domain.Comment
	for rows.Next() {
		var comment domain.Comment
		if err := rows.Scan(&comment.CommentID, &comment.TaskID, &comment.AuthorID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%s : %w", op, err)
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}

func (r *PGCommentRepository) GetById(ctx context.Context, commentID uuid.UUID) (*domain.Comment, error) {
	const op = "repository.PGCommentRepository.GetById"

	selectBuilder := sq.Select("id", "task_id", "author_id", "content", "created_at", "updated_at").
		From(comments_table).
		Where(sq.Eq{"id": commentID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	row := r.db.QueryRow(ctx, query, args...)
	var comment domain.Comment
	if err := row.Scan(&comment.CommentID, &comment.TaskID, &comment.AuthorID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrCommentNotFound
		}
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	return &comment, nil
}

func (r *PGCommentRepository) GetByTask(ctx context.Context, taskID uuid.UUID) ([]*domain.Comment, error) {
	const op = "repository.PGCommentRepository.GetByTask"

	selectBuilder := sq.Select("id", "task_id", "author_id", "content", "created_at", "updated_at").
		From(comments_table).
		Where(sq.Eq{"task_id": taskID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNoCommentsFound
		}
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	var comments []*domain.Comment
	for rows.Next() {
		var comment domain.Comment
		if err := rows.Scan(&comment.CommentID, &comment.TaskID, &comment.AuthorID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%s : %w", op, err)
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}

func (r *PGCommentRepository) Delete(ctx context.Context, commentID uuid.UUID) error {
	const op = "repository.PGCommentRepository.Delete"

	deleteBuilder := sq.Delete(comments_table).
		Where(sq.Eq{"id": commentID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ErrCommentNotFound
		}
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}

func (r *PGCommentRepository) Update(ctx context.Context, comment *domain.Comment) error {
	const op = "repository.PGCommentRepository.Update"

	updateBuilder := sq.Update(comments_table).
		Where(sq.Eq{"id": comment.CommentID}).
		Set("task_id", comment.TaskID).
		Set("author_id", comment.AuthorID).
		Set("content", comment.Content).
		Set("updated_at", comment.UpdatedAt).
		PlaceholderFormat(sq.Dollar)

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ErrCommentNotFound
		}
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}
