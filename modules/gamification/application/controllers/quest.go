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
