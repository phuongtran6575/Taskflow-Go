package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type AttachmentRepository interface {
	Create(attachment *models.Attachment) error
	Update(attachment *models.Attachment) error
	Delete(attachment *models.Attachment) error
	FindByID(id uuid.UUID) (*models.Attachment, error)
	FindAll() ([]models.Attachment, error)
}
