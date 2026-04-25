package repositories

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
	interfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces/user"
	"gorm.io/gorm"
)

type workspaceRepository struct {
	db *gorm.DB
}

func NewWorkspaceRepository(db *gorm.DB) interfaces.WorkspaceRepository {
	return &workspaceRepository{db: db}
}

func (r *workspaceRepository) Create(workspace *models.Workspace) error {
	return nil
}

func (r *workspaceRepository) Update(workspace *models.Workspace) error {
	return nil
}

func (r *workspaceRepository) Delete(workspaceID uuid.UUID) error {
	return nil
}

func (r *workspaceRepository) GetByID(id uuid.UUID) (*models.Workspace, error) {
	return nil, nil
}

func (r *workspaceRepository) GetAll() ([]models.Workspace, error) {
	return nil, nil
}
