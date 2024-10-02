package controllers

import (
    "net/http"
    "taskmanager/app/models"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
)

// Register handles user registration
func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash the password
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
    user.Password = string(hashedPassword)

    // Save the user to the database
    if err := models.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
    }

    c.JSON(http.StatusCreated, user)
}

// Login handles user authentication and token generation
func Login(c *gin.Context) {
    var user models.User
    var foundUser models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Find user by username
    if err := models.FindUserByUsername(user.Username, &foundUser); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // Check if password is correct
    if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id": foundUser.ID,
    })
    tokenString, _ := token.SignedString([]byte("your_secret_key"))

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
