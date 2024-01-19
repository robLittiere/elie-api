package gamification

import (
	controllers2 "elie-api/modules/gamification/application/controllers"
	"github.com/gin-gonic/gin"
)

func GamificationRoutes(router *gin.RouterGroup) {
	gamification := router.Group("/")
	{
		gamification.GET("/levels", controllers2.GetLevels)
		gamification.GET("/quests", controllers2.GetQuests)
		gamification.GET("/quests/user", controllers2.GetUserQuests)
	}
}
