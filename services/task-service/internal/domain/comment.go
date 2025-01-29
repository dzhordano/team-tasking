package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	CommentID uuid.UUID
	TaskID    uuid.UUID
	AuthorID  uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewComment(commentID, authorID, taskID uuid.UUID, content string) *Comment {
	return &Comment{
		CommentID: commentID,
		AuthorID:  authorID,
		TaskID:    taskID,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *Comment) Validate() error {
	if c.Content == "" {
		return errors.New("content is required")
	}
	return nil
}
