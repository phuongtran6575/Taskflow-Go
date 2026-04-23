package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type ActivityLogRepository interface {
	Create(activityLog *models.ActivityLog) error
	Update(activityLog *models.ActivityLog) error
	Delete(activityLog *models.ActivityLog) error
	GetByID(id uuid.UUID) (*models.ActivityLog, error)
	GetAll() ([]models.ActivityLog, error)
}
