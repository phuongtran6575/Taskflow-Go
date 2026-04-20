package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type ProjectRepository interface {
	Create(project *models.Project) error
	Update(project *models.Project) error
	Delete(project *models.Project) error
	FindByID(id uuid.UUID) (*models.Project, error)
	FindAll() ([]models.Project, error)
}
