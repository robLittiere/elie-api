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
	}

	quests := router.Group("/quests")
	{
		quests.GET("/", controllers.GetQuests)
		quests.GET("/daily", controllers.GetDailyQuests)
		quests.GET("/weekly", controllers.GetWeeklyQuests)
		quests.GET("/user", controllers.GetUserQuests)
		quests.POST("/user", controllers.CreateUserQuest)
		quests.PATCH("/user/progress", controllers.UpdateUserQuestProgress)
	}

	success := router.Group("/success")
	{
		success.GET("/user", controllers.GetUserSuccess)
		success.PATCH("/user/progress", controllers.UpdateUserSuccessProgress)
	}
}
