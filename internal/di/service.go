package di

import (
	"github.com/phuongtran6575/Taskflow-Go.git/internal/services"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/services/interfaces"
)

type Service struct {
	authService      interfaces.AuthService
	workspaceService interfaces.WorkspaceService
	userService      interfaces.UserService
}

func NewService(repo *Repository) *Service {
	return &Service{
		authService:      services.NewAuthService(repo.userRepo),
		workspaceService: services.NewWorkspaceService(repo.workspaceRepo),
		userService:      services.NewUserService(repo.userRepo),
	}
}
