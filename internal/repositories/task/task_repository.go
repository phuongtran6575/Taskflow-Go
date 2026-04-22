package repositories

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
	"gorm.io/gorm"
)

type taskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{DB: db}
}

func (r *taskRepository) Create(task *models.Task) error {
	result := r.DB.Create(task)
	return result.Error
}
func (r *taskRepository) Update(task *models.Task) error {
	result := r.DB.Save(task)
	return result.Error
}
func (r *taskRepository) Delete(task *models.Task) error {
	result := r.DB.Delete(task)
	return result.Error
}
func (r *taskRepository) GetByID(id uuid.UUID) (*models.Task, error) {
	var task models.Task
	result := r.DB.Where("id = ?", id).First(&task)
	return &task, result.Error
}
func (r *taskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	result := r.DB.Find(&tasks)
	return tasks, result.Error
}
