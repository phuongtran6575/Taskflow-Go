package handlers

import (
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

}
func (w *WorkspaceHandler) GetWorkspaces(c *gin.Context) {

}
func (w *WorkspaceHandler) GetWorkspace(c *gin.Context) {

}
func (w *WorkspaceHandler) DeleteWorkspace(c *gin.Context) {

}
func (w *WorkspaceHandler) UpdateWorkspace(c *gin.Context) {

}
func (w *WorkspaceHandler) CreateWorkspace(c *gin.Context) {

}
