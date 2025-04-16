package main

import (
	"log"

	"github.com/MalikSaddique/go_learning/jwt-auth-go/analyzer"
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

	r.POST("/signup", handler.SignUp)
	r.POST("/login", handler.HandleLogin)
	r.GET("/protected", handler.ProtectedHandler)
	r.POST("/refresh", handler.HandleRefresh)

	// Posting Task
	r.POST("/analyze", func(c *gin.Context) {
		result, err := analyzer.AnalyzeFile("Dummy_text.txt")
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to read file"})
			return
		}
		c.JSON(200, result)
	})

	// Start the server
	r.Run(":8000")
}
