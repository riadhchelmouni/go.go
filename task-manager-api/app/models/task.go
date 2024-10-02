package models

import (
    "database/sql"
    "errors"
    "log"

    _ "github.com/lib/pq" // PostgreSQL driver
    "golang.org/x/crypto/bcrypt"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB(dataSourceName string) error {
    var err error
    db, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        return err
    }

    // Check if the connection works
    if err = db.Ping(); err != nil {
        return err
    }

    log.Println("Database connected successfully!")
    return nil
}

// Task represents a task model
type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}

// GetAllTasks retrieves all tasks from the database
func GetAllTasks() ([]Task, error) {
    rows, err := db.Query("SELECT id, title, description, completed FROM tasks")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []Task
    for rows.Next() {
        var task Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    // Handle any errors encountered during iteration
    if err = rows.Err(); err != nil {
        return nil, err
    }

    return tasks, nil
}

// CreateTask adds a new task to the database
func CreateTask(task Task) error {
    _, err := db.Exec("INSERT INTO tasks (title, description, completed) VALUES ($1, $2, $3)",
        task.Title, task.Description, task.Completed)
    return err
}

// UpdateTask modifies an existing task
func UpdateTask(id int, task Task) error {
    _, err := db.Exec("UPDATE tasks SET title=$1, description=$2, completed=$3 WHERE id=$4",
        task.Title, task.Description, task.Completed, id)
    return err
}

// DeleteTask removes a task from the database
func DeleteTask(id int) error {
    _, err := db.Exec("DELETE FROM tasks WHERE id=$1", id)
    return err
}

// User represents a user model for authentication
type User struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"` // This should be the hashed password
}

// CreateUser adds a new user to the database
func CreateUser(user *User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    
    _, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, hashedPassword)
    return err
}

// FindUserByUsername finds a user by their username
func FindUserByUsername(username string, user *User) error {
    err := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).
        Scan(&user.ID, &user.Username, &user.Password)

    if err == sql.ErrNoRows {
        return errors.New("user not found")
    }
    return err
}

// VerifyPassword checks if the provided password matches the hashed password
func VerifyPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// CloseDB closes the database connection
func CloseDB() error {
    if db != nil {
        return db.Close()
    }
    return nil
}

