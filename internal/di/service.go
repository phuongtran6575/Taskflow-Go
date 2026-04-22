package di

import (
	"github.com/phuongtran6575/Taskflow-Go.git/internal/services"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/services/interfaces"
)

type Service struct {
	authService      interfaces.AuthService
	workspaceService interfaces.WorkspaceService
}

func NewService(repo *Repository) *Service {
	return &Service{
		authService:      services.NewAuthService(),
		workspaceService: services.NewWorkspaceService(repo.workspaceRepo),
	}
}
