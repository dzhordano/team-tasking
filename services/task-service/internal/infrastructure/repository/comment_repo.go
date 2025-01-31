package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain/repository"
	"github.com/google/uuid"
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

func (r *PGCommentRepository) List(ctx context.Context, limit, offset uint64) ([]*domain.Comment, error) {
	return nil, nil
}

func (r *PGCommentRepository) GetById(ctx context.Context, commentID uuid.UUID) (*domain.Comment, error) {
	return nil, nil
}

func (r *PGCommentRepository) GetByTask(ctx context.Context, taskID uuid.UUID) ([]*domain.Comment, error) {
	return nil, nil
}

func (r *PGCommentRepository) Delete(ctx context.Context, commentID uuid.UUID) error {
	return nil
}

func (r *PGCommentRepository) Update(ctx context.Context, comment *domain.Comment) error {
	return nil
}
