package mapper

import (
	"github.com/phuongtran6575/Taskflow-Go.git/internal/dto"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

func ToWorkspaceFromCreateWorkspaceDTO(workspaceDTO *dto.CreateWorkspaceDTO) *models.Workspace {
	return &models.Workspace{
		Name:   workspaceDTO.Name,
		Domain: workspaceDTO.Domain,
	}
}

func ToWorkspaceFromUpdateWorkspaceDTO(workspaceDTO *dto.UpdateWorkspaceDTO) *models.Workspace {
	return &models.Workspace{
		Name:   workspaceDTO.Name,
		Domain: workspaceDTO.Domain,
	}
}
