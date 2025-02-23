package config

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
	"os"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
    var err error

	dbPath := os.Getenv("DATABASE_PATH")
    if dbPath == "" {
        dbPath = "./orders.db"  // Default to the local DB file if not set in Docker
    }

    DB, err = gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB.AutoMigrate(&Order{}) // Ensure the Order model is migrated
}