package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type PermissionRepository interface {
	Create(permission *models.Permission) error
	Update(permission *models.Permission) error
	Delete(permission *models.Permission) error
	GetByID(id uuid.UUID) (*models.Permission, error)
	GetAll() ([]models.Permission, error)
}
