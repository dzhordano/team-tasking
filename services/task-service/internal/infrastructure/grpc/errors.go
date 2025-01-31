package grpc

import (
	"context"
	"errors"

	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapError(ctx context.Context, err error) error {
	unwrappedErr := errors.Unwrap(err)
	if unwrappedErr != nil {
		err = unwrappedErr
	}

	switch {
	case errors.Is(err, domain.ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, domain.ErrInvalidUUID):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, domain.ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, err.Error())

	// Task Errors
	case errors.Is(err, domain.ErrTaskNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, domain.ErrNoTasksFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, domain.ErrTaskAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())

	// Task Request Errors
	case errors.Is(err, domain.ErrNoPendingTasksFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, domain.ErrTaskRequestAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())

	// Project Errors
	case errors.Is(err, domain.ErrProjectNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, domain.ErrNoProjectsFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, domain.ErrProjectAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())

	// Comment Errors
	case errors.Is(err, domain.ErrNoCommentsFound):
		return status.Error(codes.NotFound, err.Error())

	default:
		return status.Error(codes.Internal, "internal server error")
	}
}
