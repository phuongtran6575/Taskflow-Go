package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
	Update(comment *models.Comment) error
	Delete(comment *models.Comment) error
	GetByID(id uuid.UUID) (*models.Comment, error)
	GetAll() ([]models.Comment, error)
}
