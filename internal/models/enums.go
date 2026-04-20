package models

// AuthProvider enum
type AuthProvider int

const (
	Email AuthProvider = iota
	Google
	Github
)

// WorkspacePlan enum
type WorkspacePlan int

const (
	Free WorkspacePlan = iota
	Pro
	Enterprise
)

// WorkspaceRole enum
type WorkspaceRole int

const (
	Owner WorkspaceRole = iota
	Admin
	Member
)

// TaskPriority enum
type TaskPriority int

const (
	Low TaskPriority = iota
	Medium
	High
	Urgent
)

// ActivityAction enum
type ActivityAction int

const (
	Create ActivityAction = iota
	Update
	Delete
)

// EntityType enum
type EntityType string

const (
	EntityTask      EntityType = "TASK"
	EntityProject   EntityType = "PROJECT"
	EntityComment   EntityType = "COMMENT"
	EntityColumn    EntityType = "COLUMN"
	EntityWorkspace EntityType = "WORKSPACE"
)

// NotificationType enum
type NotificationType string

const (
	NotifAssigned       NotificationType = "ASSIGNED"
	NotifMentioned      NotificationType = "MENTIONED"
	NotifCommented      NotificationType = "COMMENTED"
	NotifStatusChanged  NotificationType = "STATUS_CHANGED"
	NotifAddedToProject NotificationType = "ADDED_TO_PROJECT"
	NotifAddedToWksp    NotificationType = "ADDED_TO_WORKSPACE"
	NotifTaskDueSoon    NotificationType = "TASK_DUE_SOON"
	NotifAnnouncement   NotificationType = "ANNOUNCEMENT"
	NotifCustom         NotificationType = "CUSTOM"
)
