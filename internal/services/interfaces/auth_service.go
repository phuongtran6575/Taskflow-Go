package interfaces

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/dto"
)

type AuthService interface {
	GetProfile(id uuid.UUID) (dto.UserProfileDTO, error)
	Register(user *dto.RegisterDTO) error
	Login(dto *dto.LoginDTO) (dto.ResponseTokenDTO, error)
	Logout(id uuid.UUID) error
	ForgotPassword(email string) error
	ResetPassword(token, password string) error
}
