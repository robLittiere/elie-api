package gamification

import (
	"elie-api/modules/gamification/controllers"
	"github.com/gin-gonic/gin"
)

func GamificationRoutes(router *gin.RouterGroup) {
	gamification := router.Group("/")
	{
		gamification.GET("/levels", controllers.GetLevels)
		gamification.GET("/quests", controllers.GetQuests)
	}
}
