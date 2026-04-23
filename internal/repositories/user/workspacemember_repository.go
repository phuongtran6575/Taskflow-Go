package repositories

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
	interfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces/user"
	"gorm.io/gorm"
)

type workspaceMemberRepository struct {
	DB *gorm.DB
}

func NewWorkspaceMemberRepository(db *gorm.DB) interfaces.WorkspaceMemberRepository {
	return &workspaceMemberRepository{DB: db}
}

func (r *workspaceMemberRepository) Create(workspaceMember *models.WorkspaceMember) error {
	return nil
}

func (r *workspaceMemberRepository) Update(workspaceMember *models.WorkspaceMember) error {
	return nil
}

func (r *workspaceMemberRepository) Delete(workspaceMember *models.WorkspaceMember) error {
	return nil
}

func (r *workspaceMemberRepository) FindByUserID(id uuid.UUID) ([]models.WorkspaceMember, error) {
	return nil, nil
}

func (r *workspaceMemberRepository) FindByWorkspaceID(id uuid.UUID) ([]models.WorkspaceMember, error) {
	return nil, nil
}

func (r *workspaceMemberRepository) FindByWorkspaceIDAndUserID(workspaceID, userID uuid.UUID) (*models.WorkspaceMember, error) {
	return nil, nil
}
