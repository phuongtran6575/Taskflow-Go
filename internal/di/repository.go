package di

import (
	interfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces/user"
	repositories "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/user"
	"gorm.io/gorm"
)

type Repository struct {
	userRepo      interfaces.UserRepository
	workspaceRepo interfaces.WorkspaceRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		userRepo:      repositories.NewUserRepository(db),
		workspaceRepo: repositories.NewWorkspaceRepository(db),
	}
}
