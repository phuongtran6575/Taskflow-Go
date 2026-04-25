package dto

import "github.com/google/uuid"

type ResponseTokenDTO struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type UserProfileDTO struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	FullName     string    `json:"full_name"`
	AvatarURL    string    `json:"avatar_url"`
	AuthProvider any       `json:"auth_provider"`
}

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterDTO struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	FullName    string `json:"full_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
}

type GoogleOAuthDTO struct {
	TokenID string `json:"token_id" validate:"required"`
}

type GoogleOAuthResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}
