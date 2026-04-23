package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type WorkspaceInviteRepository interface {
	Create(workspaceInvite *models.WorkspaceInvite) error
	Update(workspaceInvite *models.WorkspaceInvite) error
	Delete(workspaceInvite *models.WorkspaceInvite) error
	FindByID(id uuid.UUID) (*models.WorkspaceInvite, error)
	FindAll() ([]models.WorkspaceInvite, error)
}
