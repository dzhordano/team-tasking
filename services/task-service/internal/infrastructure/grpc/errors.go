package grpc

import (
	"context"
	"errors"

	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapError(ctx context.Context, err error) error {
	err = errors.Unwrap(err)

	switch {
	case errors.Is(err, domain.ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, domain.ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, err.Error())
	case errors.Is(err, domain.ErrTaskNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, domain.ErrTaskAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())
	default:
		return status.Error(codes.Internal, "internal server error")
	}
}
