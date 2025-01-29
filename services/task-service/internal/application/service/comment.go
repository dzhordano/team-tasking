package services

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/application/interfaces"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain/repository"
	"github.com/google/uuid"
)

type commentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(commentRepository repository.CommentRepository) interfaces.CommentService {
	return &commentService{
		commentRepository: commentRepository,
	}
}

func (s *commentService) CreateComment(ctx context.Context, taskID uuid.UUID, authorID uuid.UUID, content string) error {
	panic("TODO")
}

func (s *commentService) GetUserComments(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*domain.Comment, error) {
	panic("TODO")
}

func (s *commentService) GetUserTaskComments(ctx context.Context, userID uuid.UUID, taskID uuid.UUID) ([]*domain.Comment, error) {
	panic("TODO")
}

func (s *commentService) UpdateComment(ctx context.Context, content string, commentID uuid.UUID) error {
	panic("TODO")
}

func (s *commentService) DeleteComment(ctx context.Context, commentID uuid.UUID) error {
	panic("TODO")
}
