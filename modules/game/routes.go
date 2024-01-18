package game

import (
	"elie-api/modules/game/application/controllers"
	"github.com/gin-gonic/gin"
)

func GameRoutes(router *gin.RouterGroup) {
	game := router.Group("/games")
	{
		game.GET("/", controllers.GetGames)
		game.GET("/quiz", controllers.GetQuizGames)
		game.POST("/", controllers.CreateGame)
	}

	quiz := router.Group("/quiz")
	{
		quiz.GET("/", controllers.GetQuizzes)
	}
}
