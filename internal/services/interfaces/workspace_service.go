package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/dto"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type WorkspaceService interface {
	Create(workspace *dto.CreateWorkspaceDTO) error
	Update(workspace *dto.UpdateWorkspaceDTO) error
	Delete(workspaceID uuid.UUID) error
	GetByID(id uuid.UUID) (*models.Workspace, error)
	GetAll() ([]models.Workspace, error)
}
