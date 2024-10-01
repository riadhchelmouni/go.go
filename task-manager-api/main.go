package main

import (
	"fmt"
	"log"
	"taskmanager/config"
	"taskmanager/app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize the database
	config.ConnectDatabase()

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	} else {
		fmt.Println("Server running on port 8080")
	}
}
