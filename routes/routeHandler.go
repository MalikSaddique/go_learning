package routes

import (
	handler "github.com/MalikSaddique/go_learning/handlers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RoutesHandler(r *gin.Engine) {
	// r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.POST("/signup", handler.SignUp)
	r.POST("/login")
	r.GET("/protected", handler.ProtectedHandler)
	r.GET("/refresh", handler.HandleRefresh)
	// r.GET("/getdata/:user_id", handler.GetResults)
}
