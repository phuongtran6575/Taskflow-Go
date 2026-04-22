package repositories

import (
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{DB: db}
}

// Create implements [interfaces.UserRepository].
func (u *userRepository) Create(user *models.User) error {
	result := u.DB.Create(user)
	return result.Error
}

// Delete implements [interfaces.UserRepository].
func (u *userRepository) Delete(user *models.User) error {
	result := u.DB.Delete(user)
	return result.Error
}

// FindAll implements [interfaces.UserRepository].
func (u *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	result := u.DB.Find(&users)
	return users, result.Error
}

// FindByEmail implements [interfaces.UserRepository].
func (u *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := u.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}

// FindByID implements [interfaces.UserRepository].
func (u *userRepository) FindByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	result := u.DB.Where("id = ?", id).First(&user)
	return &user, result.Error
}

// Update implements [interfaces.UserRepository].
func (u *userRepository) Update(user *models.User) error {
	result := u.DB.Save(user)
	return result.Error
}
