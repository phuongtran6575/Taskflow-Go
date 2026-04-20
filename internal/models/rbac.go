package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Permission model for permissions table
type Permission struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Slug        string         `gorm:"type:varchar(50);not null;unique"`
	Module      string         `gorm:"type:varchar(50);not null"`
	Description string         `gorm:"type:text"`
	IsSystem    bool           `gorm:"type:boolean;not null;default:true"`
	CreatedAt   time.Time      `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Relationships
	Roles []Role `gorm:"many2many:role_permissions"`
}

// Role model for roles table
type Role struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	WorkspaceID uuid.UUID      `gorm:"type:uuid;not null"`
	Name        string         `gorm:"type:varchar(50);not null"`
	Description string         `gorm:"type:text"`
	UpdatedAt   time.Time      `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Relationships
	Workspace   Workspace    `gorm:"foreignKey:WorkspaceID"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

// RolePermission model for role_permissions table (Composite PK)
type RolePermission struct {
	RoleID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	PermissionID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// Relationships
	Role       Role       `gorm:"foreignKey:RoleID"`
	Permission Permission `gorm:"foreignKey:PermissionID"`
}
