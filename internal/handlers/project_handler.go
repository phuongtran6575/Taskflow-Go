package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/services/interfaces"
)

type projectHandler struct {
	projectService interfaces.ProjectService
}

func NewProjectHandler(projectService interfaces.ProjectService) *projectHandler {
	return &projectHandler{projectService: projectService}
}
func (h *projectHandler) GetProjects(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetProjects"})
}
func (h *projectHandler) GetProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetProject"})
}
func (h *projectHandler) DeleteProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteProject"})
}
func (h *projectHandler) UpdateProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateProject"})
}
func (h *projectHandler) CreateProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "CreateProject"})
}
