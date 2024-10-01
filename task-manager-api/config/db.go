package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Replace the password with the correct one you set in step 1
	dsn := "host=localhost user=postgres password=new_password dbname=taskmanager port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	fmt.Println("Database connected!")
	DB = database
}
