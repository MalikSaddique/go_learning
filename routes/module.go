package routes

import (
	authservice "github.com/MalikSaddique/go_learning/service/auth_service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine      *gin.Engine
	AuthService authservice.AuthService
}
