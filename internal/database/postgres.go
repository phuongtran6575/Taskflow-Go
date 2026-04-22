package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDB initializes the PostgreSQL connection using DNS from environment variables
func ConnectDB() (*gorm.DB, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	if sslmode == "" {
		sslmode = "disable"
	}
	if timezone == "" {
		timezone = "UTC"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db
	log.Println("Database connection established")
	return db, nil
}

// Migrate runs AutoMigrate for all defined models
func Migrate(db *gorm.DB) error {
	// Ensure uuid-ossp extension is enabled for uuid_generate_v4()
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	err := db.AutoMigrate(
		&models.User{},
		&models.Workspace{},
		&models.WorkspaceMember{},
		&models.WorkspaceInvite{},
		&models.Permission{},
		&models.Role{},
		&models.RolePermission{},
		&models.Project{},
		&models.ProjectMember{},
		&models.Column{},
		&models.Task{},
		&models.TaskAssignee{},
		&models.Attachment{},
		&models.Comment{},
		&models.Label{},
		&models.TaskLabel{},
		&models.ActivityLog{},
		&models.Notification{},
		&models.NotificationRecipient{},
	)

	if err != nil {
		return fmt.Errorf("failed to run database migration: %w", err)
	}

	log.Println("Database migration completed successfully")
	return nil
}
