package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "taskmanager/app/models"
)

// HomeHandler responds with a welcome message
func HomeHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Task Manager API!"})
}

// GetTasks retrieves all tasks from the database
func GetTasks(c *gin.Context) {
    tasks, err := models.GetAllTasks() // Capture both values
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

// CreateTask adds a new task to the database
func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := models.CreateTask(task)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully!"})
}

// UpdateTask modifies an existing task
func UpdateTask(c *gin.Context) {
    id := c.Param("id") // Get the id from URL parameters
    taskID, err := strconv.Atoi(id) // Convert string to int
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = models.UpdateTask(taskID, task) // Pass the converted id
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully!"})
}

// DeleteTask removes a task from the database
func DeleteTask(c *gin.Context) {
    id := c.Param("id") // Get the id from URL parameters
    taskID, err := strconv.Atoi(id) // Convert string to int
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    err = models.DeleteTask(taskID) // Pass the converted id
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully!"})
}
