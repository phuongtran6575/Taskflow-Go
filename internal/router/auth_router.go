package router

import (
	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/handlers"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/middleware"
)

func RegisterAuthRoutes(router *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	a := router.Group("/auth")
	a.GET("/profile", middleware.AuthMiddleware(), authHandler.GetProfile)
	a.POST("/register", authHandler.Register)
	a.POST("/login", authHandler.Login)
	a.POST("/logout", authHandler.Logout)
	a.POST("/forgot-password", authHandler.ForgotPassword)
	a.POST("/reset-password", authHandler.ResetPassword)

}
