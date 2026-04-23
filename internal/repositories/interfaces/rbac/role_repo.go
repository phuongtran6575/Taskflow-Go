package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type RoleRepository interface {
	Create(role *models.Role) error
	Update(role *models.Role) error
	Delete(role *models.Role) error
	GetByID(id uuid.UUID) (*models.Role, error)
	GetAll() ([]models.Role, error)
}
