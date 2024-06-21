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
		quests.GET("/userQuests", controllers.GetUserQuests)
		quests.POST("/userQuest", controllers.CreateUserQuest)
		quests.PATCH("/userQuest/progress", controllers.UpdateUserQuestProgress)
	}

	success := router.Group("/successes")
	{
		success.GET("/", controllers.GetSuccesses)
		success.GET("/user", controllers.GetUserSuccesses)
		success.PATCH("/user/progress", controllers.UpdateUserSuccessProgress)
	}
}
