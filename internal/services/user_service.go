package services

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
	repoInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces/user"
	serviceInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/services/interfaces"
)

type userService struct {
	userRepository repoInterfaces.UserRepository
}

func NewUserService(userRepository repoInterfaces.UserRepository) serviceInterfaces.UserService {
	return &userService{userRepository: userRepository}
}
func (s *userService) Create(user *models.User) error {
	return s.userRepository.Create(user)
}
func (s *userService) Update(user *models.User) error {
	return s.userRepository.Update(user)
}
func (s *userService) Delete(user *models.User) error {
	return s.userRepository.Delete(user)
}
func (s *userService) FindByID(id uuid.UUID) (*models.User, error) {
	return s.userRepository.FindByID(id)
}
func (s *userService) FindAll() ([]models.User, error) {
	return s.userRepository.FindAll()
}
func (s *userService) FindByEmail(email string) (*models.User, error) {
	return s.userRepository.FindByEmail(email)
}
