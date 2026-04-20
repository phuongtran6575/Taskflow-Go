package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type WorkspaceRepository interface {
	Create(workspace *models.Workspace) error
	Update(workspace *models.Workspace) error
	Delete(workspace *models.Workspace) error
	FindByID(id uuid.UUID) (*models.Workspace, error)
	FindAll() ([]models.Workspace, error)
}
