package router

import (
	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/handlers"
)

func RegisterWorkspaceRoutes(router *gin.RouterGroup, workspaceHandler *handlers.WorkspaceHandler) {
	w := router.Group("/workspace")
	w.POST("/create", workspaceHandler.CreateWorkspace)
	w.DELETE("/:id", workspaceHandler.DeleteWorkspace)
	w.PUT("/:id", workspaceHandler.UpdateWorkspace)
	w.GET("/:id", workspaceHandler.GetWorkspaceByID)
	w.GET("/", workspaceHandler.GetWorkspaces)
	w.POST("/add-user", workspaceHandler.AddUserToWorkspace)
}
