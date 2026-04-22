package main

import (
	"log"

	//"github.com/phuongtran6575/Taskflow-Go.git/internal/configs"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/database"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/di"
)

func main() {
	// Load configuration
	//config := configs.LoadConfig()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	router := di.SetupDI(db)
	/*migraion := database.Migrate(db)
	if migraion != nil {
		log.Fatal("Failed to migrate database:", migraion)
	}*/

	router.Run(":8080")
}
