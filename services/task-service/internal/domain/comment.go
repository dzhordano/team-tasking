package domain

import (
	"fmt"
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

func (c *Comment) SetContent(content string) {
	c.Content = content
	c.UpdatedAt = time.Now()
}

func NewComment(commentID, taskID, authorID uuid.UUID, content string) *Comment {
	return &Comment{
		CommentID: commentID,
		TaskID:    taskID,
		AuthorID:  authorID,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *Comment) Validate() error {
	if c.Content == "" {
		return fmt.Errorf("%s : %w", "content is required", ErrInvalidArgument)
	}
	return nil
}
