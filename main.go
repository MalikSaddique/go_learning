package main

import (
	"fmt"
	"log"

	_ "github.com/MalikSaddique/go_learning/docs"
	connection "github.com/MalikSaddique/go_learning/jwt-auth-go/database"
	"github.com/MalikSaddique/go_learning/jwt-auth-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// @title File Analyzer APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @in header
// @name token
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8002
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	Router := gin.Default()

	connection.DbConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}
	fmt.Println("Databse connected successfully")

	routes.RoutesHandler(Router)

	// Start the server
	Router.Run(":8002")
}
