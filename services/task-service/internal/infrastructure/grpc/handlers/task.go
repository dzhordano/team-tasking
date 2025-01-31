package handlers

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/application/interfaces"
	"github.com/dzhordano/team-tasking/services/tasks/pkg/context/keys"
	task_v1 "github.com/dzhordano/team-tasking/services/tasks/pkg/grpc/task/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	projectId, err := uuid.Parse(req.GetProjectId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	deadline := req.GetDeadline().AsTime()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := h.ts.CreateTask(ctx, req.GetTitle(), req.GetDescription(), userId, projectId, deadline); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *TaskHandler) AssignTask(ctx context.Context, req *task_v1.AssignTaskRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (h *TaskHandler) GetUserTasks(ctx context.Context, req *task_v1.GetUserTasksRequest) (*task_v1.GetUserTasksResponse, error) {
	return &task_v1.GetUserTasksResponse{}, nil
}

func (h *TaskHandler) GetProjectTasks(ctx context.Context, req *task_v1.GetProjectTasksRequest) (*task_v1.GetProjectTasksResponse, error) {
	return &task_v1.GetProjectTasksResponse{}, nil
}

func (h *TaskHandler) GetTask(ctx context.Context, req *task_v1.GetTaskRequest) (*task_v1.GetTaskResponse, error) {
	return &task_v1.GetTaskResponse{}, nil
}

func (h *TaskHandler) UpdateTask(ctx context.Context, req *task_v1.UpdateTaskRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (h *TaskHandler) DeleteTask(ctx context.Context, req *task_v1.DeleteTaskRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
