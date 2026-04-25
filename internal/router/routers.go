package router

import (
	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/handlers"
)

type Router struct {
	authHandler      *handlers.AuthHandler
	workspaceHandler *handlers.WorkspaceHandler
	userHandler      *handlers.UserHandler
}

func NewRouter(authHandler *handlers.AuthHandler, workspaceHandler *handlers.WorkspaceHandler, userHandler *handlers.UserHandler) *Router {
	return &Router{
		authHandler:      authHandler,
		workspaceHandler: workspaceHandler,
		userHandler:      userHandler,
	}
}

func (r *Router) RegisterRoutes() *gin.Engine {
	engine := gin.Default()

	api := engine.Group("/api/v1")

	RegisterAuthRoutes(api, r.authHandler)
	RegisterWorkspaceRoutes(api, r.workspaceHandler)
	return engine
}
