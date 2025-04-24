package authservice

import (
	"github.com/MalikSaddique/go_learning/models"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(c *gin.Context, login *models.UserLogin) (*models.TokenPair, error)
}
