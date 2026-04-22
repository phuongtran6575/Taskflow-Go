package di

import "github.com/phuongtran6575/Taskflow-Go.git/internal/handlers"

type handler struct {
	authHandler      *handlers.AuthHandler
	workspaceHandler *handlers.WorkspaceHandler
}

func NewHandler(s *Service) *handler {
	return &handler{
		authHandler:      handlers.NewAuthHandler(s.authService),
		workspaceHandler: handlers.NewWorkspaceHandler(s.workspaceService),
	}
}
