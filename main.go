package main

import (
	"elie-api/config"
	"elie-api/modules/auth"
	"elie-api/modules/game"
	"elie-api/modules/gamification"
	"elie-api/modules/user"
	"elie-api/modules/websocket"
	"elie-api/modules/websocket/dualquiz"
	"elie-api/modules/websocket/matchmaking"
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

	// Setup websocket matchHub
	matchHub := matchmaking.NewHub()
	dqHub := dualquiz.NewDqHub()
	matchHub.WithRoomCreator(dqHub)

	hubs := map[string]websocket.Hub{
		"matchmaking": matchHub,
		"dualquiz":    dqHub,
	}

	go matchHub.Run()
	go dqHub.Run()

	// Serve static index html for websocket tests
	router.Static("/static", "./static")

	api := router.Group("/api")
	version := api.Group("/v1")

	auth.AuthRoutes(version)
	user.UserRoutes(version)
	gamification.GamificationRoutes(version)
	game.GameRoutes(version)
	websocket.WebsocketRoutes(router, hubs)

	return router
}
