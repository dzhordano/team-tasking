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
	commentID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	comment := domain.NewComment(commentID, taskID, authorID, content)

	if err := comment.Validate(); err != nil {
		return err
	}

	if err := s.commentRepository.Save(ctx, comment); err != nil {
		return err
	}

	return nil
}

func (s *commentService) GetUserComments(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*domain.Comment, error) {
	comments, err := s.commentRepository.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *commentService) GetUserTaskComments(ctx context.Context, userID uuid.UUID, taskID uuid.UUID) ([]*domain.Comment, error) {
	comments, err := s.commentRepository.GetByTask(ctx, taskID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *commentService) UpdateComment(ctx context.Context, content string, commentID uuid.UUID) error {
	comment, err := s.commentRepository.GetById(ctx, commentID)
	if err != nil {
		return err
	}

	comment.SetContent(content)

	if err := s.commentRepository.Update(ctx, comment); err != nil {
		return err
	}

	return nil
}

func (s *commentService) DeleteComment(ctx context.Context, commentID uuid.UUID) error {
	if err := s.commentRepository.Delete(ctx, commentID); err != nil {
		return err
	}

	return nil
}
