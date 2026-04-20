package models

import (
	"time"

	"github.com/google/uuid"
)

// Notification model for notifications table
type Notification struct {
	ID           uuid.UUID        `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ActorID      *uuid.UUID       `gorm:"type:uuid"`
	Type         NotificationType `gorm:"type:varchar(50);not null"`
	Title        string           `gorm:"type:varchar(255);not null"`
	Content      string           `gorm:"type:text"`
	ReferenceURL string           `gorm:"type:text"`
	CreatedAt    time.Time        `gorm:"type:timestamp;not null;default:now()"`

	// Relationships
	Actor      *User                   `gorm:"foreignKey:ActorID"`
	Recipients []NotificationRecipient `gorm:"foreignKey:NotificationID"`
}

// NotificationRecipient model for notification_recipients table (Composite PK)
type NotificationRecipient struct {
	NotificationID uuid.UUID `gorm:"type:uuid;primaryKey"`
	RecipientID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	IsRead         bool      `gorm:"type:boolean;not null;default:false"`
	ReadAt         *time.Time `gorm:"type:timestamp"`

	// Relationships
	Notification Notification `gorm:"foreignKey:NotificationID"`
	Recipient    User         `gorm:"foreignKey:RecipientID"`
}
