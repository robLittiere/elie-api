package game

import (
	"elie-api/modules/game/application/controllers"
	"github.com/gin-gonic/gin"
)

func GameRoutes(router *gin.RouterGroup) {
	game := router.Group("/games")
	{
		game.GET("/", controllers.GetGames)
		game.POST("/", controllers.CreateGame)
	}
}
