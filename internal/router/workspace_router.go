package router

import (
	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/handlers"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/middleware"
)

func RegisterWorkspaceRoutes(router *gin.RouterGroup, workspaceHandler *handlers.WorkspaceHandler) {
	w := router.Group("/workspace")
	w.Use(middleware.AuthMiddleware())
	w.POST("/create", workspaceHandler.CreateWorkspace)
	w.DELETE("/:id", workspaceHandler.DeleteWorkspace)
	w.PUT("/:id", workspaceHandler.UpdateWorkspace)
	w.GET("/:id", workspaceHandler.GetWorkspaceByID)
	w.GET("/", workspaceHandler.GetWorkspaces)
	w.POST("/add-user", workspaceHandler.AddUserToWorkspace)
}
