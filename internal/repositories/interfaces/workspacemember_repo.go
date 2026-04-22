package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type WorkspaceMemberRepository interface {
	Create(member *models.WorkspaceMember) error
	FindByUserID(userID uuid.UUID) ([]models.WorkspaceMember, error)
	FindByWorkspaceID(workspaceID uuid.UUID) ([]models.WorkspaceMember, error)
	FindByWorkspaceIDAndUserID(workspaceID, userID uuid.UUID) (*models.WorkspaceMember, error)
	Update(member *models.WorkspaceMember) error
	Delete(member *models.WorkspaceMember) error
	DeleteByWorkspaceIDAndUserID(workspaceID, userID uuid.UUID) error
}
