package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type TaskLabelRepository interface {
	Create(taskLabel *models.TaskLabel) error
	Update(taskLabel *models.TaskLabel) error
	Delete(taskLabel *models.TaskLabel) error
	GetByID(id uuid.UUID) (*models.TaskLabel, error)
	GetAll() ([]models.TaskLabel, error)
}
