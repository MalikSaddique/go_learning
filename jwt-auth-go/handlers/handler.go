package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/MalikSaddique/go_learning/jwt-auth-go/analyzer"
	"github.com/MalikSaddique/go_learning/jwt-auth-go/auth"
	connection "github.com/MalikSaddique/go_learning/jwt-auth-go/database"
	"github.com/MalikSaddique/go_learning/jwt-auth-go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")
var refreshSecretKey = []byte("my_refresh_secret_key")

type UserInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//handle login

func HandleLogin(c *gin.Context) {
	var u UserInfo
	var dbEmail, dbPassword string
	db := connection.DbConnection()

	err := json.NewDecoder(c.Request.Body).Decode(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON token"})
	}

	query := `SELECT email, password FROM users WHERE email = $1`
	err = db.QueryRow(query, u.Email).Scan(&dbEmail, &dbPassword)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if dbPassword != u.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	tokenString, err := auth.CreateToken(u.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token not generated"})
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": u.Email,
			"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
		})

	refreshTokenString, err := refreshToken.SignedString([]byte(refreshSecretKey))
	if err != nil {
		fmt.Printf("Refresh tocken not generated")
	}
	c.JSON(http.StatusOK, gin.H{
		"Access Token Key":  tokenString,
		"Refresh Token key": refreshTokenString,
	})

}

//protected handler

func ProtectedHandler(c *gin.Context) {
	db := connection.DbConnection()

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	err := auth.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}
	result, err := analyzer.AnalyzeFile("Dummy_text.txt")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read file"})
		return
	}

	err = connection.SaveResult(db, result)
	if err != nil {
		log.Fatal("Failed to save the results")
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data saved successfully",
	})
}

//refresh token handler

func HandleRefresh(c *gin.Context) {
	refreshTokenString := c.GetHeader("Authorization")

	refreshTokenString = strings.TrimPrefix(refreshTokenString, "Bearer ")

	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return []byte(refreshSecretKey), nil
	})

	if err != nil || !refreshToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	if claims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
		email := claims["email"].(string)

		newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Minute * 15).Unix(),
		})

		newAccessTokenString, err := newAccessToken.SignedString([]byte(secretKey))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new access token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token": newAccessTokenString,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
	}

}

//signup handling

func SignUp(c *gin.Context) {
	var user models.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	db := connection.DbConnection()

	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING email`
	err = db.QueryRow(query, user.Email, user.Password).Scan(&user.Email, &user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User already existed with this email"})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "User created successfully",
		"email":   user.Email,
	})
}

//result handling
