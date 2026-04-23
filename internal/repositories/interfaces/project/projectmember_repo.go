package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type ProjectMemberRepository interface {
	Create(projectMember *models.ProjectMember) error
	Update(projectMember *models.ProjectMember) error
	Delete(projectMember *models.ProjectMember) error
	GetByID(id uuid.UUID) (*models.ProjectMember, error)
	GetAll() ([]models.ProjectMember, error)
}
