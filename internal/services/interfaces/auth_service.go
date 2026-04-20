package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

type AuthService interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
	Logout(id uuid.UUID) error
	ForgotPassword(email string) error
	ResetPassword(token, password string) error
}
