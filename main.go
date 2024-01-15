package main

import (
	"elie-api/cmd/api/config"
	"elie-api/cmd/api/modules/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDb()
	config.Migrate()

	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	api := router.Group("/api")
	version := api.Group("/v1")

	auth.AuthRoutes(version)
	return router
}
