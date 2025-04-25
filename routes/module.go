package routes

import (
	authservice "github.com/MalikSaddique/go_learning/service/auth_service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine      *gin.Engine
	AuthService authservice.AuthService
}

func NewRouter(authService authservice.AuthService) *Router {
	engine := gin.Default()
	router := &Router{
		Engine:      engine,
		AuthService: authService,
	}
	router.defineRoutes()
	return router
}
