package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq" // PostgreSQL driver
    "taskmanager/app/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Database connection string
    dsn := "host=localhost user=postgres password=new_password dbname=taskmanager port=5432 sslmode=disable"
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create the tasks table
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        description TEXT,
        completed BOOLEAN DEFAULT FALSE
    )`)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Database migrated successfully!")

    // Initialize the Gin router
    r := gin.Default()

    // Setup routes
    routes.SetupRoutes(r)

    // Start the server
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Could not start server: %v\n", err)
    } else {
        fmt.Println("Server running on port 8080")
    }
}
