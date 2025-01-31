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

type TaskHandler struct {
	task_v1.UnimplementedTaskServiceV1Server
	ts interfaces.TaskService
}

func NewTaskHandler(ts interfaces.TaskService) *TaskHandler {
	return &TaskHandler{
		ts: ts,
	}
}

func (h *TaskHandler) CreateTask(ctx context.Context, req *task_v1.CreateTaskRequest) (*emptypb.Empty, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	projectId, err := uuid.Parse(req.GetProjectId())
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	deadline := req.GetDeadline().AsTime()

	if err := h.ts.CreateTask(ctx, req.GetTitle(), req.GetDescription(), userId, projectId, deadline); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *TaskHandler) AssignTask(ctx context.Context, req *task_v1.AssignTaskRequest) (*emptypb.Empty, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	taskId, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	assigneeId, err := uuid.Parse(req.GetAssigneeId())
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	if err := h.ts.AssignTask(ctx, userId, taskId, assigneeId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *TaskHandler) GetUserTasks(ctx context.Context, req *task_v1.GetUserTasksRequest) (*task_v1.GetUserTasksResponse, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	tasks, err := h.ts.GetUserTasks(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &task_v1.GetUserTasksResponse{
		Tasks: converter.TasksToGRPC(tasks),
	}, nil
}

func (h *TaskHandler) GetUserPendingTasks(ctx context.Context, req *task_v1.GetUserPendingTasksRequest) (*task_v1.GetUserPendingTasksResponse, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	tasks, err := h.ts.GetPendingTasks(ctx, userId, req.GetLimit(), req.GetOffset())
	if err != nil {
		return nil, err
	}

	return &task_v1.GetUserPendingTasksResponse{
		Tasks: converter.TasksToGRPC(tasks),
	}, nil
}

func (h *TaskHandler) AcceptTask(ctx context.Context, req *task_v1.AcceptTaskRequest) (*emptypb.Empty, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	taskId, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	if err := h.ts.AcceptTask(ctx, taskId, userId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *TaskHandler) DeclineTask(ctx context.Context, req *task_v1.DeclineTaskRequest) (*emptypb.Empty, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	taskId, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	if err := h.ts.DeclineTask(ctx, taskId, userId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *TaskHandler) FinishTask(ctx context.Context, req *task_v1.FinishTaskRequest) (*emptypb.Empty, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	taskId, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	if err := h.ts.FinishTask(ctx, taskId, userId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
