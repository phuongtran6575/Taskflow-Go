package services

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
	repoInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces"
	serInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/services/interfaces"
)

type projectService struct {
	projectRepository repoInterfaces.ProjectRepository
}

func NewProjectService(projectRepository repoInterfaces.ProjectRepository) serInterfaces.ProjectService {
	return &projectService{projectRepository: projectRepository}
}

// Create implements [interfaces.ProjectService].
func (p *projectService) Create(project *models.Project) error {
	return p.projectRepository.Create(project)
}

// Delete implements [interfaces.ProjectService].
func (p *projectService) Delete(project *models.Project) error {
	return p.projectRepository.Delete(project)
}

// FindAll implements [interfaces.ProjectService].
func (p *projectService) GetAll() ([]models.Project, error) {
	return p.projectRepository.GetAll()
}

// FindByID implements [interfaces.ProjectService].
func (p *projectService) GetByID(id uuid.UUID) (*models.Project, error) {
	return p.projectRepository.GetByID(id)
}

// Update implements [interfaces.ProjectService].
func (p *projectService) Update(project *models.Project) error {
	return p.projectRepository.Update(project)
}
