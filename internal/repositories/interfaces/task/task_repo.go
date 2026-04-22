package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type TaskRepository interface {
	Create(task *models.Task) error
	Update(task *models.Task) error
	Delete(task *models.Task) error
	GetByID(id uuid.UUID) (*models.Task, error)
	GetAll() ([]models.Task, error)
}
