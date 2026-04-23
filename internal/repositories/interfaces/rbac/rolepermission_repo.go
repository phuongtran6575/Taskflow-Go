package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type RolePermissionRepository interface {
	Create(rolePermission *models.RolePermission) error
	Update(rolePermission *models.RolePermission) error
	Delete(rolePermission *models.RolePermission) error
	GetByID(id uuid.UUID) (*models.RolePermission, error)
	GetAll() ([]models.RolePermission, error)
}
