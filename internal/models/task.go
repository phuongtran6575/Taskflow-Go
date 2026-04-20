package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Task model for tasks table
type Task struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ColumnID    uuid.UUID      `gorm:"type:uuid;not null"`
	ProjectID   uuid.UUID      `gorm:"type:uuid;not null"`
	CreatorID   *uuid.UUID     `gorm:"type:uuid"`
	ParentID    *uuid.UUID     `gorm:"type:uuid"`
	TaskNumber  int            `gorm:"type:int;not null"`
	Title       string         `gorm:"type:varchar(255);not null"`
	Description string         `gorm:"type:text"`
	Priority    TaskPriority   `gorm:"type:int;not null;default:1"`
	StartDate   *time.Time     `gorm:"type:timestamp"`
	DueDate     *time.Time     `gorm:"type:timestamp"`
	Position    float64        `gorm:"type:float;not null"`
	CreatedAt   time.Time      `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt   time.Time      `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Relationships
	Column      Column         `gorm:"foreignKey:ColumnID"`
	Project     Project        `gorm:"foreignKey:ProjectID"`
	Creator     *User          `gorm:"foreignKey:CreatorID"`
	Parent      *Task          `gorm:"foreignKey:ParentID"`
	Subtasks    []Task         `gorm:"foreignKey:ParentID"`
	Assignees   []User         `gorm:"many2many:task_assignees"`
	Attachments []Attachment   `gorm:"foreignKey:TaskID"`
	Comments    []Comment      `gorm:"foreignKey:TaskID"`
	Labels      []Label        `gorm:"many2many:task_labels"`
}

// TaskAssignee model for task_assignees table (Composite PK)
type TaskAssignee struct {
	TaskID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	AssignedAt time.Time `gorm:"type:timestamp;not null;default:now()"`

	// Relationships
	Task Task `gorm:"foreignKey:TaskID"`
	User User `gorm:"foreignKey:UserID"`
}

// Attachment model for attachments table
type Attachment struct {
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	TaskID     uuid.UUID      `gorm:"type:uuid;not null"`
	UploaderID *uuid.UUID     `gorm:"type:uuid"`
	FileName   string         `gorm:"type:varchar(255);not null"`
	FileURL    string         `gorm:"type:text;not null"`
	FileType   string         `gorm:"type:varchar(50);not null"`
	SizeBytes  int64          `gorm:"type:bigint;not null"`
	CreatedAt  time.Time      `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	// Relationships
	Task     Task  `gorm:"foreignKey:TaskID"`
	Uploader *User `gorm:"foreignKey:UploaderID"`
}

// Comment model for comments table
type Comment struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	TaskID    uuid.UUID      `gorm:"type:uuid;not null"`
	UserID    *uuid.UUID     `gorm:"type:uuid"`
	Content   string         `gorm:"type:text;not null"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationships
	Task Task  `gorm:"foreignKey:TaskID"`
	User *User `gorm:"foreignKey:UserID"`
}

// Label model for labels table (Taxonomy/Tag)
type Label struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProjectID uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"type:varchar(50);not null"`
	Color     string    `gorm:"type:varchar(7);not null"`

	// Relationships
	Project Project `gorm:"foreignKey:ProjectID"`
	Tasks   []Task  `gorm:"many2many:task_labels"`
}

// TaskLabel model for task_labels table (Composite PK)
type TaskLabel struct {
	TaskID  uuid.UUID `gorm:"type:uuid;primaryKey"`
	LabelID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// Relationships
	Task  Task  `gorm:"foreignKey:TaskID"`
	Label Label `gorm:"foreignKey:LabelID"`
}
