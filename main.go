package main

import (
	"fmt"

	handler "github.com/MalikSaddique/go_learning/jwt-auth-go/handlers"
	"github.com/gin-gonic/gin"
)




func main() {
	fmt.Println("Starting Mux server on :8000")
	r := gin.Default()

	r.POST("/login", handler.HandleLogin)
	r.GET("/protected", handler.ProtectedHandler)
	r.POST("/refresh", handler.HandleRefresh)

	fmt.Println("Starting the Gin server...")
	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Could not start the server:", err)
	}

}
