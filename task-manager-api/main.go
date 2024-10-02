package main

import (
    "taskmanager/app/controllers"
    "taskmanager/app/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Authentication routes
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    // Other routes
    routes.SetupRoutes(r)

    r.Run(":8080") // Start the server
}
