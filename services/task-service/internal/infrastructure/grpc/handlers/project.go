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

type ProjectHandler struct {
	task_v1.UnimplementedProjectServiceV1Server
	ps interfaces.ProjectService
}

func NewProjectHandler(ps interfaces.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		ps: ps,
	}
}

func (h *ProjectHandler) CreateProject(ctx context.Context, req *task_v1.CreateProjectRequest) (*emptypb.Empty, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	if err := h.ps.CreateProject(ctx, req.GetName(), userId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *ProjectHandler) GetUserProjects(ctx context.Context, req *emptypb.Empty) (*task_v1.GetUserProjectsResponse, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	projects, err := h.ps.GetUserProjects(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &task_v1.GetUserProjectsResponse{
		Projects: converter.ProjectsToGRPC(projects),
	}, nil
}

func (h *ProjectHandler) DeleteProject(ctx context.Context, req *task_v1.DeleteProjectRequest) (*emptypb.Empty, error) {
	userIdCtx := ctx.Value(keys.UserIDKey).(string)

	userId, err := uuid.Parse(userIdCtx)
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	projectId, err := uuid.Parse(req.GetProjectId())
	if err != nil {
		return nil, domain.ErrInvalidUUID
	}

	if err := h.ps.DeleteProject(ctx, userId, projectId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
