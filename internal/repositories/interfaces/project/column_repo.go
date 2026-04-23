package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type ColumnRepository interface {
	Create(column *models.Column) error
	Update(column *models.Column) error
	Delete(column *models.Column) error
	GetByID(id uuid.UUID) (*models.Column, error)
	GetAll() ([]models.Column, error)
}
