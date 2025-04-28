package userserviceimpl

import (
	"net/http"
	"strings"

	"github.com/MalikSaddique/go_learning/analyzer"
	"github.com/MalikSaddique/go_learning/auth"
	"github.com/MalikSaddique/go_learning/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")
var refreshSecretKey = []byte("my_refresh_secret_key")

func (u *UserServiceImpl) SaveResults(c *gin.Context, req *models.Result) *models.Result {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		return nil
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return nil
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID in token"})
		return nil
	}
	userID := int(userIDFloat)

	result, err := analyzer.AnalyzeFile(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to analyze file"})
		return nil
	}

	result.UserID = userID
	savedResult, err := u.UserAuth.SaveResult(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save result"})
		return nil
	}

	return savedResult
}

func (u *UserServiceImpl) GetUserResults(c *gin.Context, page int, limit int) ([]models.Result, error) {
	id := c.Param("user_id")

	offset := (page - 1) * limit

	results, err := u.UserAuth.FetchResultsByUserID(id, limit, offset)
	if err != nil {
		return nil, err
	}

	return results, nil
}
