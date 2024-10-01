package models

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

// Initialize your database connection here
func InitDB(dataSourceName string) {
    var err error
    db, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Fatal(err)
    }
}

// Task represents a task model
type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}

// GetAllTasks retrieves all tasks from the database
func GetAllTasks() []Task {
    rows, err := db.Query("SELECT id, title, description, completed FROM tasks")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var tasks []Task
    for rows.Next() {
        var task Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
            log.Fatal(err)
        }
        tasks = append(tasks, task)
    }
    return tasks
}

// CreateTask adds a new task to the database
func CreateTask(task Task) {
    _, err := db.Exec("INSERT INTO tasks (title, description, completed) VALUES ($1, $2, $3)",
        task.Title, task.Description, task.Completed)
    if err != nil {
        log.Fatal(err)
    }
}

// UpdateTask modifies an existing task
func UpdateTask(id string, task Task) {
    _, err := db.Exec("UPDATE tasks SET title=$1, description=$2, completed=$3 WHERE id=$4",
        task.Title, task.Description, task.Completed, id)
    if err != nil {
        log.Fatal(err)
    }
}

// DeleteTask removes a task
func DeleteTask(id string) {
    _, err := db.Exec("DELETE FROM tasks WHERE id=$1", id)
    if err != nil {
        log.Fatal(err)
    }
}
