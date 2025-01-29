package repository

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
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
