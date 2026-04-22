package di

import "github.com/phuongtran6575/Taskflow-Go.git/internal/handlers"

type Handler struct {
	authHandler      *handlers.AuthHandler
	workspaceHandler *handlers.WorkspaceHandler
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		authHandler:      handlers.NewAuthHandler(s.authService),
		workspaceHandler: handlers.NewWorkspaceHandler(s.workspaceService),
	}
}
