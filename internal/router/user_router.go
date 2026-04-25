package router

import (
	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/handlers"
)

func RegisterUserRoutes(router *gin.RouterGroup, handler *handlers.UserHandler) {
	router.GET("/users", handler.FindAll)
	router.GET("/users/:id", handler.FindByID)
}
