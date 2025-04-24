package routes

import (
	"net/http"

	"github.com/MalikSaddique/go_learning/models"
	"github.com/gin-gonic/gin"
)

func (r *Router) Login(c *gin.Context) {
	var req models.UserLoginReq
	var login models.UserLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	login = models.UserLogin{
		Email:    req.Email,
		Password: req.Password,
	}

	tokenPair, err := r.AuthService.Login(c, &login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokenPair)
}
