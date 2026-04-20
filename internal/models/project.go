package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Project model for projects table
type Project struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	WorkspaceID uuid.UUID      `gorm:"type:uuid;not null"`
	OwnerID     *uuid.UUID     `gorm:"type:uuid"`
	Name        string         `gorm:"type:varchar(100);not null"`
	Key         string         `gorm:"type:varchar(10);not null"`
	Icon        string         `gorm:"type:varchar(50)"`
	Background  string         `gorm:"type:varchar(255);not null;default:'#ffffff'"`
	IsArchived  bool           `gorm:"type:boolean;not null;default:false"`
	CreatedAt   time.Time      `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt   time.Time      `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Relationships
	Workspace Workspace       `gorm:"foreignKey:WorkspaceID"`
	Owner     *User           `gorm:"foreignKey:OwnerID"`
	Members   []ProjectMember `gorm:"foreignKey:ProjectID"`
	Columns   []Column        `gorm:"foreignKey:ProjectID"`
}

// ProjectMember model for project_members table (Composite PK)
type ProjectMember struct {
	ProjectID  uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	RoleID     uuid.UUID `gorm:"type:uuid;not null"`
	IsFavorite bool      `gorm:"type:boolean;not null;default:false"`
	JoinedAt   time.Time `gorm:"type:timestamp;not null;default:now()"`

	// Relationships
	Project Project `gorm:"foreignKey:ProjectID"`
	User    User    `gorm:"foreignKey:UserID"`
	Role    Role    `gorm:"foreignKey:RoleID"`
}

// Column model for columns table (Kanban/Board)
type Column struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProjectID uuid.UUID      `gorm:"type:uuid;not null"`
	Title     string         `gorm:"type:varchar(50);not null"`
	Position  float64        `gorm:"type:float;not null"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationships
	Project Project `gorm:"foreignKey:ProjectID"`
	Tasks   []Task  `gorm:"foreignKey:ColumnID"`
}
