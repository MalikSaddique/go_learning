package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/MalikSaddique/go_learning/jwt-auth-go/auth"
	connection "github.com/MalikSaddique/go_learning/jwt-auth-go/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// HandleLogin godoc
// @Summary      Login a user
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body  models.User  true  "User Credentials"
// @Success      200
// @Failure      401
// @Router       /login [post]
func HandleLogin(c *gin.Context) {
	var u UserInfo
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON token"})
		return
	}

	db, err := connection.DbConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}
	defer db.Close()

	// var user models.User
	var user UserInfo
	fmt.Println("Looking up email:", u.Email)
	row := db.QueryRow("SELECT * FROM users WHERE email = $1", u.Email)

	err = row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
	)
	fmt.Println("User ID from DB:", user.Id)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	if user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	tokenString, err := auth.CreateToken(u.Email, u.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token not generated"})
		return
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(refreshSecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Access Token Key":  tokenString,
		"Refresh Token Key": refreshTokenString,
		"Unique Id":         user.Id,
	})
}
