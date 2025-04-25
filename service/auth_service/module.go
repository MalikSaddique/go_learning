package authservice

import (
	"github.com/MalikSaddique/go_learning/models"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(c *gin.Context, login *models.UserLogin) (*models.TokenPair, error)
	SignUp(c *gin.Context, req *models.User) *models.User
}
