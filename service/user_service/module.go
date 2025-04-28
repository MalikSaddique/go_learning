package userservice

import (
	"github.com/MalikSaddique/go_learning/models"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	SaveResults(c *gin.Context, req *models.Result) *models.Result
	GetUserResults(c *gin.Context, page int, limit int) ([]models.Result, error)
}
