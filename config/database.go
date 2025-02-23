package config

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB.AutoMigrate(&Order{}) // Ensure the Order model is migrated
}