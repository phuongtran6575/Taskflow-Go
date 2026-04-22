package services

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
	serviceInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/services/interfaces"
)

type authService struct {
}

func NewAuthService() serviceInterfaces.AuthService {
	return &authService{}
}

// ForgotPassword implements [interfaces.AuthService].
func (a *authService) ForgotPassword(email string) error {
	panic("unimplemented")
}

// Login implements [interfaces.AuthService].
func (a *authService) Login(email string, password string) (*models.User, error) {
	panic("unimplemented")
}

// Logout implements [interfaces.AuthService].
func (a *authService) Logout(id uuid.UUID) error {
	panic("unimplemented")
}

// Register implements [interfaces.AuthService].
func (a *authService) Register(user *models.User) error {
	panic("unimplemented")
}

// ResetPassword implements [interfaces.AuthService].
func (a *authService) ResetPassword(token string, password string) error {
	panic("unimplemented")
}
