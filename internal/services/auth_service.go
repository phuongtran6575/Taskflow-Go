package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/dto"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/helper"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/mapper"
	repoInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces/user"
	serviceInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/services/interfaces"
)

type authService struct {
	userRepository repoInterfaces.UserRepository
}

// GetProfile implements [interfaces.AuthService].
func (a *authService) GetProfile(id uuid.UUID) (dto.UserProfileDTO, error) {
	user, err := a.userRepository.FindByID(id)
	if err != nil {
		return dto.UserProfileDTO{}, errors.New("user not found")
	}
	return mapper.ToUserProfileDTO(*user), nil
}

func NewAuthService(userRepository repoInterfaces.UserRepository) serviceInterfaces.AuthService {
	return &authService{userRepository: userRepository}
}

// ForgotPassword implements [interfaces.AuthService].
func (a *authService) ForgotPassword(email string) error {
	panic("unimplemented")
}

// Login implements [interfaces.AuthService].
func (a *authService) Login(loginDTO *dto.LoginDTO) (dto.ResponseTokenDTO, error) {
	user, err := a.userRepository.FindByEmail(loginDTO.Email)
	if err != nil {
		return dto.ResponseTokenDTO{}, err
	}

	if user == nil {
		return dto.ResponseTokenDTO{}, errors.New("user not found")
	}

	if !helper.ValidatePassword(loginDTO.Password, user.PasswordHash) {
		return dto.ResponseTokenDTO{}, errors.New("invalid password")
	}
	token, expiresIn, err := helper.GenerateJWTToken(user.Email, user.ID)
	if err != nil {
		return dto.ResponseTokenDTO{}, err
	}
	return dto.ResponseTokenDTO{
		AccessToken:  token,
		TokenType:    "Bearer",
		RefreshToken: "",
		ExpiresIn:    expiresIn,
	}, nil
}

// Logout implements [interfaces.AuthService].
func (a *authService) Logout(id uuid.UUID) error {
	panic("unimplemented")
}

// Register implements [interfaces.AuthService].
func (a *authService) Register(registerDTO *dto.RegisterDTO) error {
	user := mapper.ToUserFromRegisterDTO(*registerDTO)
	user.PasswordHash = helper.HashPassword(user.PasswordHash)
	return a.userRepository.Create(user)
}

// ResetPassword implements [interfaces.AuthService].
func (a *authService) ResetPassword(token string, password string) error {
	panic("unimplemented")
}
