package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model for users table
type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PhoneNumber  string         `gorm:"type:varchar(15);not null;unique"`
	Email        string         `gorm:"type:varchar(255);not null;unique"`
	PasswordHash string         `gorm:"type:text"`
	FullName     string         `gorm:"type:varchar(100);not null"`
	AvatarURL    string         `gorm:"type:text"`
	AuthProvider AuthProvider   `gorm:"type:int;not null;default:0"`
	IsActive     bool           `gorm:"type:boolean;not null;default:true"`
	LastLogin    *time.Time     `gorm:"type:timestamp"`
	CreatedAt    time.Time      `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt    time.Time      `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	// Relationships
	OwnedWorkspaces []Workspace       `gorm:"foreignKey:OwnerID"`
	Memberships     []WorkspaceMember `gorm:"foreignKey:UserID"`
}

// Workspace model for workspaces table
type Workspace struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	OwnerID   uuid.UUID      `gorm:"type:uuid;not null"`
	Name      string         `gorm:"type:varchar(100);not null"`
	Domain    string         `gorm:"type:varchar(100)"`
	Plan      WorkspacePlan  `gorm:"type:int;not null;default:0"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationships
	Owner   User              `gorm:"foreignKey:OwnerID"`
	Members []WorkspaceMember `gorm:"foreignKey:WorkspaceID"`
	Invites []WorkspaceInvite `gorm:"foreignKey:WorkspaceID"`
	Roles   []Role            `gorm:"foreignKey:WorkspaceID"`
}

// WorkspaceMember model for workspace_members table (Composite PK)
type WorkspaceMember struct {
	WorkspaceID uuid.UUID     `gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID     `gorm:"type:uuid;primaryKey"`
	Role        WorkspaceRole `gorm:"type:int;not null;default:2"`
	JoinedAt    time.Time     `gorm:"type:timestamp;not null;default:now()"`

	// Relationships
	Workspace Workspace `gorm:"foreignKey:WorkspaceID"`
	User      User      `gorm:"foreignKey:UserID"`
}

// WorkspaceInvite model for workspace_invites table
type WorkspaceInvite struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	WorkspaceID uuid.UUID      `gorm:"type:uuid;not null"`
	Code        string         `gorm:"type:varchar(50);not null;unique"`
	Role        string         `gorm:"type:varchar(20);not null"`
	MaxUses     *int           `gorm:"type:int"`
	UsesCount   int            `gorm:"type:int;not null;default:0"`
	CreatedBy   *uuid.UUID     `gorm:"type:uuid"`
	ExpiresAt   *time.Time     `gorm:"type:timestamp"`
	CreatedAt   time.Time      `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Relationships
	Workspace Workspace `gorm:"foreignKey:WorkspaceID"`
	Creator   *User     `gorm:"foreignKey:CreatedBy"`
}
