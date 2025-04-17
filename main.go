package main

import (
	"fmt"
	"log"

	connection "github.com/MalikSaddique/go_learning/jwt-auth-go/database"
	handler "github.com/MalikSaddique/go_learning/jwt-auth-go/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	r := gin.Default()

	connection.DbConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}

	r.POST("/signup", handler.SignUp)
	r.POST("/login", handler.HandleLogin)
	r.GET("/protected", handler.ProtectedHandler)
	r.POST("/refresh", handler.HandleRefresh)
	r.GET("/getdata/:user_id", handler.GetResults)

	// Start the server
	r.Run(":8000")
}
