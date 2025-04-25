package routes

func (r *Router) defineRoutes() {
	// r := gin.Default()
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Engine.POST("/signup", r.SignUp)
	r.Engine.POST("/login", r.Login)
	// r.GET("/protected", handler.ProtectedHandler)
	// r.GET("/refresh", handler.HandleRefresh)
	// r.GET("/getdata/:user_id", handler.GetResults)
}
