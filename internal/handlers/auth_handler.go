package handlers

import (
	"net/http"

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
	c.JSON(http.StatusOK, gin.H{"message": "Register"})
}

// Login implements [interfaces.AuthHandler].
func (a *AuthHandler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login"})
}

// Logout implements [interfaces.AuthHandler].
func (a *AuthHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout"})
}

// ForgotPassword implements [interfaces.AuthHandler].
func (a *AuthHandler) ForgotPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ForgotPassword"})
}

// ResetPassword implements [interfaces.AuthHandler].
func (a *AuthHandler) ResetPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ResetPassword"})
}
