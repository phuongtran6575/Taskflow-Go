package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/services/interfaces"
)

type AuthHandler struct {
	authService interfaces.AuthService
}

func NewAuthHandler(authService interfaces.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register implements [interfaces.AuthHandler].
func (a *AuthHandler) Register(c *gin.Context) {
}

// Login implements [interfaces.AuthHandler].
func (a *AuthHandler) Login(c *gin.Context) {
}

// Logout implements [interfaces.AuthHandler].
func (a *AuthHandler) Logout(c *gin.Context) {
}

// ForgotPassword implements [interfaces.AuthHandler].
func (a *AuthHandler) ForgotPassword(c *gin.Context) {
}

// ResetPassword implements [interfaces.AuthHandler].
func (a *AuthHandler) ResetPassword(c *gin.Context) {
}
