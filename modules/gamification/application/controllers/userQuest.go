package controllers

import (
	"elie-api/config"
	"elie-api/modules/gamification/domain/service"
	"elie-api/modules/gamification/infrastructure"
	"elie-api/modules/gamification/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserQuests(c *gin.Context) {
	userQuestRepo := infrastructure.NewUserQuestRepo(config.DB)

	queryParams := c.Request.URL.Query()

	userQuests, err := userQuestRepo.BuildQueryAndFind(queryParams)
	fmt.Printf("UserQuests: %v", userQuests)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, userQuests)
}

func CreateUserQuest(c *gin.Context) {
	userProgressReq := models.UserQuestProgressRequest{}

	if err := c.ShouldBindJSON(&userProgressReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userQuestService := service.NewUserQuestService(config.DB)
	err := userQuestService.CreateUserQuest(userProgressReq.UserUuid, userProgressReq.QuestId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "UserQuest created"})
}

func UpdateUserQuestProgress(c *gin.Context) {
	userProgressReq := models.UserQuestProgressRequest{}
	userQuestService := service.NewUserQuestService(config.DB)

	if err := c.ShouldBindJSON(&userProgressReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := userQuestService.HandleUserQuestProgress(userProgressReq.UserUuid, userProgressReq.QuestId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "UserQuest updated"})
}
