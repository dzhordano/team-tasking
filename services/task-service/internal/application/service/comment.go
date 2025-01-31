package services

import (
	"context"
	"log/slog"

	"github.com/dzhordano/team-tasking/services/tasks/internal/application/interfaces"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain/repository"
	"github.com/google/uuid"
)

type commentService struct {
	log               *slog.Logger
	commentRepository repository.CommentRepository
	taskRepository    repository.TaskRepository
}

func NewCommentService(log *slog.Logger, commentRepository repository.CommentRepository, taskRepository repository.TaskRepository) interfaces.CommentService {
	return &commentService{
		log:               log,
		commentRepository: commentRepository,
		taskRepository:    taskRepository,
	}
}

func (s *commentService) CreateComment(ctx context.Context, taskID uuid.UUID, authorID uuid.UUID, content string) error {
	_, err := s.taskRepository.GetById(ctx, taskID)
	if err != nil {
		s.log.Error("task not found", slog.String("task_id", taskID.String()))
		return err
	}

	commentID, err := uuid.NewUUID()
	if err != nil {
		s.log.Error("failed to generate comment id", slog.String("error", err.Error()))
		return err
	}

	comment := domain.NewComment(commentID, taskID, authorID, content)

	if err := comment.Validate(); err != nil {
		s.log.Error("failed to validate comment", slog.String("error", err.Error()))
		return err
	}

	if err := s.commentRepository.Save(ctx, comment); err != nil {
		s.log.Error("failed to save comment", slog.String("error", err.Error()))
		return err
	}

	s.log.Debug("comment created", slog.String("comment_id", commentID.String()))

	return nil
}

func (s *commentService) GetUserComments(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*domain.Comment, error) {
	if limit == 0 {
		limit = 10 // FIXME magic number
	}

	comments, err := s.commentRepository.GetUserComments(ctx, userID, limit, offset)
	if err != nil {
		s.log.Error("failed to get user comments", slog.String("error", err.Error()))
		return nil, err
	}

	if len(comments) == 0 {
		s.log.Debug("no comments found for user", slog.String("user_id", userID.String()))
		return nil, domain.ErrNoCommentsFound
	}

	s.log.Debug("user comments found", slog.String("user_id", userID.String()))

	return comments, nil
}

func (s *commentService) GetUserTaskComments(ctx context.Context, userID uuid.UUID, taskID uuid.UUID) ([]*domain.Comment, error) {
	comments, err := s.commentRepository.GetByTask(ctx, taskID)
	if err != nil {
		s.log.Error("failed to get task comments", slog.String("error", err.Error()))
		return nil, err
	}

	if len(comments) == 0 {
		s.log.Debug("no comments found for task", slog.String("task_id", taskID.String()))
		return nil, domain.ErrNoCommentsFound
	}

	s.log.Debug("task comments found", slog.String("task_id", taskID.String()))

	return comments, nil
}

func (s *commentService) UpdateComment(ctx context.Context, content string, commentID, userID uuid.UUID) error {
	comment, err := s.commentRepository.GetById(ctx, commentID)
	if err != nil {
		s.log.Error("comment not found", slog.String("comment_id", commentID.String()))
		return err
	}

	if comment.AuthorID != userID {
		s.log.Error("permission denied", slog.String("comment_id", commentID.String()))
		return domain.ErrCommentNotFound
	}

	comment.SetContent(content)

	if err := s.commentRepository.Update(ctx, comment); err != nil {
		s.log.Error("failed to update comment", slog.String("error", err.Error()))
		return err
	}

	s.log.Debug("comment updated", slog.String("comment_id", commentID.String()))

	return nil
}

func (s *commentService) DeleteComment(ctx context.Context, commentID, userID uuid.UUID) error {
	comment, err := s.commentRepository.GetById(ctx, commentID)
	if err != nil {
		s.log.Error("comment not found", slog.String("comment_id", commentID.String()))
		return err
	}

	if comment.AuthorID != userID {
		s.log.Error("permission denied", slog.String("comment_id", commentID.String()))
		return domain.ErrCommentNotFound
	}

	if err := s.commentRepository.Delete(ctx, commentID); err != nil {
		s.log.Error("failed to delete comment", slog.String("error", err.Error()))
		return err
	}

	s.log.Debug("comment deleted", slog.String("comment_id", commentID.String()))

	return nil
}
