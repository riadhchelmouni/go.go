package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "taskmanager/app/models"
)

// HomeHandler serves as a basic home endpoint
func HomeHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Task Manager API"})
}

// GetTasks retrieves all tasks
func GetTasks(c *gin.Context) {
    tasks := models.GetAllTasks() // You need to implement this function in the models package
    c.JSON(http.StatusOK, tasks)
}

// CreateTask adds a new task
func CreateTask(c *gin.Context) {
    var task models.Task // Define your Task model accordingly
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    models.CreateTask(task) // Implement this in models
    c.JSON(http.StatusCreated, task)
}

// UpdateTask modifies an existing task
func UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    models.UpdateTask(id, task) // Implement this in models
    c.JSON(http.StatusOK, task)
}

// DeleteTask removes a task
func DeleteTask(c *gin.Context) {
    id := c.Param("id")
    models.DeleteTask(id) // Implement this in models
    c.JSON(http.StatusNoContent, nil)
}
