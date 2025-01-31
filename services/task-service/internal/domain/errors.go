package domain

import "errors"

var (
	ErrInvalidArgument  = errors.New("invalid argument")
	ErrInvalidUUID      = errors.New("invalid uuid")
	ErrPermissionDenied = errors.New("permission denied")

	ErrTaskNotFound      = errors.New("task not found")
	ErrNoTasksFound      = errors.New("no tasks found")
	ErrTaskAlreadyExists = errors.New("task already exists")

	ErrNoPendingTasksFound      = errors.New("no pending tasks found")
	ErrTaskRequestAlreadyExists = errors.New("task request already exists")

	ErrProjectNotFound      = errors.New("project not found")
	ErrNoProjectsFound      = errors.New("no projects found")
	ErrProjectAlreadyExists = errors.New("project already exists")

	ErrNoCommentsFound = errors.New("no comments found")
	ErrCommentNotFound = errors.New("comment not found")
)
