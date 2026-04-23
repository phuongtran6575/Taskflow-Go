package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type LabelRepository interface {
	Create(label *models.Label) error
	Update(label *models.Label) error
	Delete(label *models.Label) error
	GetByID(id uuid.UUID) (*models.Label, error)
	GetAll() ([]models.Label, error)
}
