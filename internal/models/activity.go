package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// ActivityLog model for activity_logs table
type ActivityLog struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	WorkspaceID *uuid.UUID     `gorm:"type:uuid"`
	ProjectID   *uuid.UUID     `gorm:"type:uuid"`
	UserID      *uuid.UUID     `gorm:"type:uuid"`
	Action      ActivityAction `gorm:"type:varchar(50);not null"`
	EntityType  EntityType     `gorm:"type:varchar(20);not null"`
	EntityID    uuid.UUID      `gorm:"type:uuid;not null"`
	Metadata    datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt   time.Time      `gorm:"type:timestamp;not null;default:now()"`

	// Relationships
	Workspace *Workspace `gorm:"foreignKey:WorkspaceID"`
	Project   *Project   `gorm:"foreignKey:ProjectID"`
	User      *User      `gorm:"foreignKey:UserID"`
}
