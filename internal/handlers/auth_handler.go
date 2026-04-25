package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/dto"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/services/interfaces"
)

type AuthHandler struct {
	authService interfaces.AuthService
}

func NewAuthHandler(authService interfaces.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (a *AuthHandler) GetProfile(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		dto.Fail(c, http.StatusUnauthorized, "UNAUTHORIZED", "User ID not found in token")
		return
	}
	user, err := a.authService.GetProfile(userId.(uuid.UUID))
	if err != nil {
		dto.Fail(c, http.StatusUnauthorized, "UNAUTHORIZED", err.Error())
		return
	}
	dto.OK(c, user)
}

// Register implements [interfaces.AuthHandler].
func (a *AuthHandler) Register(c *gin.Context) {
	var register dto.RegisterDTO
	if err := c.ShouldBindJSON(&register); err != nil {
		dto.Fail(c, http.StatusBadRequest, "INVALID_INPUT", err.Error())
		return
	}

	// Validate DTO
	/*if err := utils.ValidateStruct(&register); err != nil {
		dto.Fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}*/

	if err := a.authService.Register(&register); err != nil {
		dto.Fail(c, http.StatusInternalServerError, "REGISTER_FAILED", err.Error())
		return
	}

	dto.OK(c, register)
}

// Login implements [interfaces.AuthHandler].
func (a *AuthHandler) Login(c *gin.Context) {
	var login dto.LoginDTO
	if err := c.ShouldBindJSON(&login); err != nil {
		dto.Fail(c, http.StatusBadRequest, "INVALID_INPUT", err.Error())
		return
	}

	token, err := a.authService.Login(&login)
	if err != nil {
		dto.Fail(c, http.StatusInternalServerError, "LOGIN_FAILED", err.Error())
		return
	}

	/*if err := utils.ValidateStruct(&login); err != nil {
		dto.Fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}*/

	dto.OK(c, token)
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
