package di

import (
	activityInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces/activity"
	projectInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces/project"
	rbacInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces/rbac"
	taskInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces/task"
	userInterfaces "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/interfaces/user"
	userRepositories "github.com/phuongtran6575/Taskflow-Go.git/internal/repositories/user"
	"gorm.io/gorm"
)

type Repository struct {
	userRepo            userInterfaces.UserRepository
	workspaceRepo       userInterfaces.WorkspaceRepository
	projectRepo         projectInterfaces.ProjectRepository
	workspaceMemberRepo userInterfaces.WorkspaceMemberRepository
	workspaceInviteRepo userInterfaces.WorkspaceInviteRepository
	activityLogRepo     activityInterfaces.ActivityLogRepository
	attachmentRepo      taskInterfaces.AttachmentRepository
	labelRepo           taskInterfaces.LabelRepository
	taskLabelRepo       taskInterfaces.TaskLabelRepository
	taskRepo            taskInterfaces.TaskRepository
	commentRepo         taskInterfaces.CommentRepository
	permissionRepo      rbacInterfaces.PermissionRepository
	rolePermissionRepo  rbacInterfaces.RolePermissionRepository
	roleRepo            rbacInterfaces.RoleRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		userRepo:            userRepositories.NewUserRepository(db),
		workspaceRepo:       userRepositories.NewWorkspaceRepository(db),
		workspaceMemberRepo: userRepositories.NewWorkspaceMemberRepository(db),
		workspaceInviteRepo: userRepositories.NewWorkspaceInviteRepository(db),
	}
}
