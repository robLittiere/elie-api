package controllers

import (
	"elie-api/config"
	"elie-api/modules/gamification/infrastructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetQuests(c *gin.Context) {
	questRepo := infrastructure.NewQuestRepo(config.DB)
	queryParams := c.Request.URL.Query()

	quests, err := questRepo.BuildQueryAndFind(queryParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, quests)
}

func GetDailyQuests(context *gin.Context) {
	questRepo := infrastructure.NewQuestRepo(config.DB)
	quests, err := questRepo.BuildQueryAndFind(map[string][]string{"type": {"daily"}})
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// TODO: Add logic to get the 3 daily quests

	context.JSON(200, quests)
}

func GetWeeklyQuests(context *gin.Context) {
	questRepo := infrastructure.NewQuestRepo(config.DB)
	quests, err := questRepo.BuildQueryAndFind(map[string][]string{"type": {"weekly"}})
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(200, quests)
}
