package successes

import (
	"elie-api/modules/gamification/application/controllers"
	updateUserSuccess "elie-api/modules/gamification/successes/UpdateUserSuccess"
	"elie-api/modules/gamification/successes/listSuccesses"
	"elie-api/modules/gamification/successes/listUserSuccesses"
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
		success.GET("/", listSuccesses.GetSuccesses)
		success.GET("/user", listUserSuccesses.GetUserSuccesses)
		success.PATCH("/user/progress", updateUserSuccess.UpdateUserSuccessProgress)
	}
}
