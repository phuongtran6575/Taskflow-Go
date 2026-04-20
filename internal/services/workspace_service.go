package services

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
	repoInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces"
	serviceInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/services/interfaces"
)

type workspaceService struct {
	workspaceRepository repoInterfaces.WorkspaceRepository
}

func NewWorkspaceService(workspaceRepository repoInterfaces.WorkspaceRepository) serviceInterfaces.WorkspaceService {
	return &workspaceService{workspaceRepository: workspaceRepository}
}

func (s *workspaceService) Create(workspace *models.Workspace) error {
	return s.workspaceRepository.Create(workspace)
}

func (s *workspaceService) Update(workspace *models.Workspace) error {
	return s.workspaceRepository.Update(workspace)
}

func (s *workspaceService) Delete(workspace *models.Workspace) error {
	return s.workspaceRepository.Delete(workspace)
}

func (s *workspaceService) FindByID(id uuid.UUID) (*models.Workspace, error) {
	return s.workspaceRepository.FindByID(id)
}

func (s *workspaceService) FindAll() ([]models.Workspace, error) {
	return s.workspaceRepository.FindAll()
}
