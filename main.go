package main

import (
	"elie-api/config"
	"elie-api/modules/auth"
	"elie-api/modules/game"
	"elie-api/modules/gamification"
	"elie-api/modules/user"
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
	user.UserRoutes(version)
	gamification.GamificationRoutes(version)
	game.GameRoutes(version)
	return router
}
