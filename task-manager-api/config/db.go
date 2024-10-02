package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "taskmanager/app/models"
)

// InitDB initializes the database connection from environment variables
func InitDB() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Get database connection string from environment variables
    dsn := os.Getenv("DATABASE_URL")

    // Initialize the database connection
    err = models.InitDB(dsn)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    log.Println("Database connection successful!")
}
