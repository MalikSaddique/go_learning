package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// HandleRefresh godoc
// @Summary      Refresh Access Token
// @Description  Validates refresh token and generates a new access token
// @Tags         auth
// @Security     BearerAuth
// @Produce      json
// @Success      200  "Returns new access token"
// @Failure      401  "Unauthorized"
// @Failure      500 "Internal Server Error"
// @Router       /refresh [get]
func HandleRefresh(c *gin.Context) {
	refreshTokenString := c.GetHeader("Authorization")

	refreshTokenString = strings.TrimPrefix(refreshTokenString, "Bearer ")

	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(refreshSecretKey), nil

	})
	if err != nil {
		log.Printf("JWT parse error: %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	if err != nil || !refreshToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}
	if refreshTokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is empty"})
		return
	}

	if claims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
		email := claims["email"].(string)
		// id := claims["user_id"].(int)
		// var userid UserInfo

		newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": email,
			// "user_id": userid.Id,
			"exp": time.Now().Add(time.Minute * 15).Unix(),
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
