package di

import (
	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/router"
	"gorm.io/gorm"
)

func SetupDI(db *gorm.DB) *gin.Engine {
	repositories := NewRepository(db)
	services := NewService(repositories)
	handlers := NewHandler(services)

	router := router.NewRouter(
		handlers.authHandler,
		handlers.workspaceHandler,
		handlers.userHandler,
	)
	return router.RegisterRoutes()
}
