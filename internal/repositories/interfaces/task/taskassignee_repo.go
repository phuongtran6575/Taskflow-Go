package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type TaskAssigneeRepository interface {
	Create(taskAssignee *models.TaskAssignee) error
	Update(taskAssignee *models.TaskAssignee) error
	Delete(taskAssignee *models.TaskAssignee) error
	GetByID(id uuid.UUID) (*models.TaskAssignee, error)
	GetAll() ([]models.TaskAssignee, error)
}
