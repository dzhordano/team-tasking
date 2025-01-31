package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ProjectID uuid.UUID
	OwnerID   uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProject(projectID, ownerID uuid.UUID, name string) *Project {
	return &Project{
		ProjectID: projectID,
		OwnerID:   ownerID,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (p *Project) Validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
