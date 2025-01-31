package interfaces

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/google/uuid"
)

type CommentService interface {
	CreateComment(ctx context.Context, taskID uuid.UUID, authorID uuid.UUID, content string) error
	GetUserComments(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*domain.Comment, error)
	GetUserTaskComments(ctx context.Context, userID uuid.UUID, taskID uuid.UUID) ([]*domain.Comment, error)
	UpdateComment(ctx context.Context, content string, commentID, userID uuid.UUID) error
	DeleteComment(ctx context.Context, commentID, userID uuid.UUID) error
}
