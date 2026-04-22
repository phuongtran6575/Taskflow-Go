package services

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
	repoInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces/task"
	serviceInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/services/interfaces"
)

type taskService struct {
	taskRepo repoInterfaces.TaskRepository
}

func NewTaskService(taskRepo repoInterfaces.TaskRepository) serviceInterfaces.TaskService {
	return &taskService{taskRepo: taskRepo}
}
func (s *taskService) Create(task *models.Task) error {
	return s.taskRepo.Create(task)
}
func (s *taskService) Update(task *models.Task) error {
	return s.taskRepo.Update(task)
}
func (s *taskService) Delete(task *models.Task) error {
	return s.taskRepo.Delete(task)
}
func (s *taskService) GetByID(id uuid.UUID) (*models.Task, error) {
	return s.taskRepo.GetByID(id)
}
func (s *taskService) GetAll() ([]models.Task, error) {
	return s.taskRepo.GetAll()
}
