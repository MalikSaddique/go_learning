package main

import (
	"log"

	"github.com/MalikSaddique/go_learning/database"
	_ "github.com/MalikSaddique/go_learning/docs"
	"github.com/MalikSaddique/go_learning/routes"
	authserviceimpl "github.com/MalikSaddique/go_learning/service/auth_service/auth_service_impl"
	userserviceimpl "github.com/MalikSaddique/go_learning/service/user_service/user_service_impl"
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

	conn, err := database.DbConnection()

	userdb := database.NewStorage(conn)

	authService := authserviceimpl.NewAuthService(authserviceimpl.NewAuthServiceImpl{
		UserAuth: userdb,
	})
	userService := userserviceimpl.NewUserService(userdb)

	router := routes.NewRouter(authService, userService)

	router.Engine.Run(":8002")

}
