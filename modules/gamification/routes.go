package gamification

import (
	"elie-api/modules/gamification/application/controllers"
	"github.com/gin-gonic/gin"
)

func GamificationRoutes(router *gin.RouterGroup) {
	gamification := router.Group("/")
	{
		gamification.GET("/levels", controllers.GetLevels)
		gamification.GET("/quests", controllers.GetQuests)
		gamification.GET("/quests/user", controllers.GetUserQuests)
		gamification.GET("/quests/daily", controllers.GetDailyQuests)
		gamification.GET("/quests/weekly", controllers.GetWeeklyQuests)
	}
}
