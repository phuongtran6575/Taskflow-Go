package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/services/interfaces"
)

type WorkspaceHandler struct {
	workspaceService interfaces.WorkspaceService
}

func NewWorkspaceHandler(workspaceService interfaces.WorkspaceService) *WorkspaceHandler {
	return &WorkspaceHandler{workspaceService: workspaceService}
}

func (w *WorkspaceHandler) AddUserToWorkspace(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddUserToWorkspace"})
}
func (w *WorkspaceHandler) GetWorkspaces(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetWorkspaces"})
}
func (w *WorkspaceHandler) GetWorkspaceByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "GetWorkspace", "id": id})
}
func (w *WorkspaceHandler) DeleteWorkspace(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteWorkspace"})
}
func (w *WorkspaceHandler) UpdateWorkspace(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateWorkspace"})
}
func (w *WorkspaceHandler) CreateWorkspace(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "CreateWorkspace"})
}
