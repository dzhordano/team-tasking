package domain

import "errors"

var (
	ErrInvalidArgument  = errors.New("invalid argument")
	ErrInvalidUUID      = errors.New("invalid uuid")
	ErrPermissionDenied = errors.New("permission denied")

	ErrTaskNotFound      = errors.New("task not found")
	ErrTaskAlreadyExists = errors.New("task already exists")

	ErrProjectNotFound      = errors.New("project not found")
	ErrProjectAlreadyExists = errors.New("project already exists")
)
