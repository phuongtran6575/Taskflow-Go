package repositories

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type projectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) interfaces.ProjectRepository {
	return &projectRepository{DB: db}
}

// Create implements [interfaces.ProjectRepository].
func (p *projectRepository) Create(project *models.Project) error {
	result := p.DB.Create(project)
	return result.Error
}

// Delete implements [interfaces.ProjectRepository].
func (p *projectRepository) Delete(project *models.Project) error {
	result := p.DB.Delete(project)
	return result.Error
}

// FindAll implements [interfaces.ProjectRepository].
func (p *projectRepository) GetAll() ([]models.Project, error) {
	var projects []models.Project
	result := p.DB.Find(&projects)
	return projects, result.Error
}

// FindByID implements [interfaces.ProjectRepository].
func (p *projectRepository) GetByID(id uuid.UUID) (*models.Project, error) {
	var project models.Project
	result := p.DB.Where("id = ?", id).First(&project)
	return &project, result.Error
}

// Update implements [interfaces.ProjectRepository].
func (p *projectRepository) Update(project *models.Project) error {
	result := p.DB.Save(project)
	return result.Error
}
