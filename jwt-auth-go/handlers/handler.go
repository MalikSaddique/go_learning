package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/MalikSaddique/go_learning/jwt-auth-go/analyzer"
	"github.com/MalikSaddique/go_learning/jwt-auth-go/auth"
	connection "github.com/MalikSaddique/go_learning/jwt-auth-go/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")
var refreshSecretKey = []byte("my_refresh_secret_key")

type UserInfo struct {
	Id       int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//protected handler

func ProtectedHandler(c *gin.Context) {

	db, err := connection.DbConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	userIDFloat, ok := claims["user_id"].(float64)
	fmt.Printf("Token claims: %+v\n", claims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID in token"})
		return
	}
	userID := int(userIDFloat)

	result, err := analyzer.AnalyzeFile("Dummy_text.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to analyze file"})
		return
	}

	result.ID = userID

	err = connection.SaveResult(db, result)
	if err != nil {
		log.Println("Failed to save the results:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save result"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data saved successfully",
	})

}
