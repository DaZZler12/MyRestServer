package main

import (
	"context"
	"log"
	"time"

	"github.com/DaZZler12/MyRestServer/pkg/config"
	"github.com/DaZZler12/MyRestServer/pkg/database"
	"github.com/DaZZler12/MyRestServer/pkg/handlers"
	"github.com/DaZZler12/MyRestServer/pkg/middlewares"
	"github.com/DaZZler12/MyRestServer/pkg/privateroutes"
	"github.com/DaZZler12/MyRestServer/pkg/publicroutes"
	"github.com/DaZZler12/MyRestServer/pkg/service"
	"github.com/DaZZler12/MyRestServer/pkg/store"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func init() {
	configFilePath := "../config/master.yaml"
	_, err := config.ReadConfig(configFilePath)
	if err != nil {
		log.Fatal("Error on App-Startup: ", err)
	}
	storeInstance := store.GetStore()
	userService := service.NewUserService(storeInstance)
	handlers.SetUserService(userService)
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	server := gin.Default()
	server.Use(apmgin.Middleware(server))
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-BrowserFingerprint", "X-Workspace-ID"},
		ExposeHeaders:    []string{"Content-Length", "Content-Range", "X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	defer database.DisconnectMongoDB(context.Background()) // disconnect from mongo if application shutdown.
	public := server.Group("/api")
	h := handlers.New()
	publicroutes.PublicRoutes(public, h)
	privateRoutes := server.Group("/api/items")
	privateRoutes.Use(middlewares.JwtAuthMiddleware())
	privateroutes.Privateroutes(privateRoutes, h)
	log.Fatal(server.Run(":8080"))
}
