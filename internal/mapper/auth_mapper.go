package mapper

import (
	"github.com/phuongtran6575/Taskflow-Go.git/internal/dto"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
)

func ToUserFromRegisterDTO(dto dto.RegisterDTO) *models.User {
	return &models.User{
		Email:        dto.Email,
		PasswordHash: dto.Password,
		FullName:     dto.FullName,
		PhoneNumber:  dto.PhoneNumber,
	}
}
func ToUserProfileDTO(user models.User) dto.UserProfileDTO {
	return dto.UserProfileDTO{
		ID:           user.ID,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		FullName:     user.FullName,
		AvatarURL:    user.AvatarURL,
		AuthProvider: user.AuthProvider,
	}
}
