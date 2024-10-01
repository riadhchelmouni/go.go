package routes

import (
    "github.com/gin-gonic/gin"
    "taskmanager/app/controllers"
)

func SetupRoutes(router *gin.Engine) {
    router.GET("/", controllers.HomeHandler)
    router.GET("/tasks", controllers.GetTasks)
    router.POST("/tasks", controllers.CreateTask)
    router.PUT("/tasks/:id", controllers.UpdateTask)
    router.DELETE("/tasks/:id", controllers.DeleteTask)
}
