package repositories

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
	interfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces/user"
	"gorm.io/gorm"
)

type workspaceInviteRepository struct {
	db *gorm.DB
}

func NewWorkspaceInviteRepository(db *gorm.DB) interfaces.WorkspaceInviteRepository {
	return &workspaceInviteRepository{db: db}
}

func (r *workspaceInviteRepository) Create(workspaceInvite *models.WorkspaceInvite) error {
	return nil
}

func (r *workspaceInviteRepository) Update(workspaceInvite *models.WorkspaceInvite) error {
	return nil
}

func (r *workspaceInviteRepository) Delete(workspaceInvite *models.WorkspaceInvite) error {
	return nil
}

func (r *workspaceInviteRepository) FindByID(id uuid.UUID) (*models.WorkspaceInvite, error) {
	return nil, nil
}

func (r *workspaceInviteRepository) FindAll() ([]models.WorkspaceInvite, error) {
	return nil, nil
}
