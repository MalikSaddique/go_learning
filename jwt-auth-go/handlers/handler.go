package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/MalikSaddique/go_learning/jwt-auth-go/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")
var refreshSecretKey = []byte("my_refresh_secret_key")
var hardCodedEmail = "email@123.com"
var hardCodedPassword = "12345"

type UserInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleLogin(c *gin.Context) {
	var u UserInfo
	// w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(c.Request.Body).Decode(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON token"})
	}

	if u.Email != hardCodedEmail || u.Password != hardCodedPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized Access"})
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

func ProtectedHandler(c *gin.Context) {
	// w.Header().Set("Content-Type", "application/json")
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
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome! You are authorized",
	})
}


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
