package handlers

import (
	"context"

	"github.com/dzhordano/team-tasking/services/tasks/internal/application/interfaces"
	task_v1 "github.com/dzhordano/team-tasking/services/tasks/pkg/grpc/task/v1"
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
	return &emptypb.Empty{}, nil
}

func (h *ProjectHandler) GetUserProjects(ctx context.Context, req *task_v1.GetUserProjectsRequest) (*task_v1.GetUserProjectsResponse, error) {
	return &task_v1.GetUserProjectsResponse{}, nil
}

func (h *ProjectHandler) DeleteProject(ctx context.Context, req *task_v1.DeleteProjectRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
