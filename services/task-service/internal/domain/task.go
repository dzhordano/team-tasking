package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	NilUUID = uuid.Nil
)

type Task struct {
	TaskID      uuid.UUID
	ProjectID   uuid.UUID
	AssigneeID  uuid.UUID
	Title       string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deadline    time.Time
}

type TaskStatus string

const (
	TaskTODO       TaskStatus = "TODO"
	TaskINPROGRESS TaskStatus = "IN_PROGRESS"
	TaskDONE       TaskStatus = "DONE"
	TaskARCHIVED   TaskStatus = "ARCHIVED"
)

func (t *Task) ChangeStatus(newStatus TaskStatus) error {
	if t.Status == TaskTODO && newStatus == TaskDONE {
		return errors.New("cannot change status from TODO to DONE")
	}

	t.Status = newStatus
	t.UpdatedAt = time.Now()

	return nil
}

func (t *Task) SetAssignee(assigneeID uuid.UUID) {
	t.AssigneeID = assigneeID
	t.UpdatedAt = time.Now()
}

func NewTask(taskID, projectID uuid.UUID, title, description string, deadline time.Time) *Task {
	return &Task{
		TaskID:      taskID,
		ProjectID:   projectID,
		AssigneeID:  NilUUID, // Because we will have to assign someone for task afterwards.
		Title:       title,
		Description: description,
		Status:      TaskTODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Deadline:    deadline,
	}
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	if t.Description == "" {
		return errors.New("description is required")
	}
	return nil
}
