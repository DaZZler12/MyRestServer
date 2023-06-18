package main

import (
	"context"
	"log"
	"net/http"
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
	cfg, err := config.ReadConfig(configFilePath)
	if err != nil {
		panic(err)
	}

	storeInstance := store.GetStore(cfg.Database)
	userService := service.NewUserService(storeInstance)
	handlers.SetUserService(userService)
}

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
	server.GET("/swagger/*any", gin.WrapH(http.FileServer(http.Dir("./docs"))))
	public := server.Group("/api")
	h := handlers.New()
	publicroutes.PublicRoutes(public, h)
	privateRoutes := server.Group("/api/items")
	privateRoutes.Use(middlewares.JwtAuthMiddleware())
	privateroutes.Privateroutes(privateRoutes, h)
	log.Fatal(server.Run(":8080"))
}
