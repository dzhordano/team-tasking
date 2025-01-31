package handlers

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/application/interfaces"
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"github.com/dzhordano/team-tasking/services/tasks/internal/infrastructure/grpc/converter"
	"github.com/dzhordano/team-tasking/services/tasks/pkg/context/keys"
	task_v1 "github.com/dzhordano/team-tasking/services/tasks/pkg/grpc/task/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommentHandler struct {
	task_v1.UnimplementedCommentServiceV1Server
	cs interfaces.CommentService
}

func NewCommentService(cs interfaces.CommentService) *CommentHandler {
	return &CommentHandler{
		cs: cs,
	}
}

func (h *CommentHandler) CreateComment(ctx context.Context, req *task_v1.CreateCommentRequest) (*emptypb.Empty, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	taskId, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	if err := h.cs.CreateComment(ctx, taskId, userId, req.GetContent()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *CommentHandler) GetUserComments(ctx context.Context, req *task_v1.GetUserCommentsRequest) (*task_v1.GetUserCommentsResponse, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	comments, err := h.cs.GetUserComments(ctx, userId, req.GetLimit(), req.GetOffset())
	if err != nil {
		return nil, err
	}

	return &task_v1.GetUserCommentsResponse{
		Comments: converter.CommentsToGRPC(comments),
	}, nil
}

func (h *CommentHandler) GetUserTaskComments(ctx context.Context, req *task_v1.GetUserTaskCommentsRequest) (*task_v1.GetUserTaskCommentsResponse, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	taskId, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	comments, err := h.cs.GetUserTaskComments(ctx, userId, taskId)
	if err != nil {
		return nil, err
	}

	return &task_v1.GetUserTaskCommentsResponse{
		Comments: converter.CommentsToGRPC(comments),
	}, nil
}

func (h *CommentHandler) UpdateComment(ctx context.Context, req *task_v1.UpdateCommentRequest) (*emptypb.Empty, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	commentId, err := uuid.Parse(req.GetCommentId())
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	if err := h.cs.UpdateComment(ctx, req.GetContent(), commentId, userId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *CommentHandler) DeleteComment(ctx context.Context, req *task_v1.DeleteCommentRequest) (*emptypb.Empty, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	commentId, err := uuid.Parse(req.GetCommentId())
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	if err := h.cs.DeleteComment(ctx, commentId, userId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
