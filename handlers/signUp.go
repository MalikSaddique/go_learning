package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	connection "github.com/MalikSaddique/go_learning/database"
	"github.com/MalikSaddique/go_learning/models"
	"github.com/gin-gonic/gin"
)

// SignUp godoc
// @Summary      Register a new user
// @Description  Create a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body  models.User  true  "User Info"
// @Success      201
// @Failure      400
// @Router       /signup [post]
func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db, err := connection.DbConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	var existingUser models.User
	err = db.QueryRow("SELECT email FROM users WHERE email = $1", user.Email).Scan(&existingUser.Email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists with this email"})
		return
	} else if err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking existing user"})
		return
	}

	_, err = db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"email":   user.Email,
	})
}
