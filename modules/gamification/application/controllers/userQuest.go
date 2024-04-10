package controllers

import (
	"elie-api/config"
	"elie-api/modules/gamification/infrastructure"
	"elie-api/modules/gamification/models"
	infraUser "elie-api/modules/user/infrastructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserQuests(c *gin.Context) {
	userQuestRepo := infrastructure.NewUserQuestRepo(config.DB)

	queryParams := c.Request.URL.Query()

	userQuests, err := userQuestRepo.BuildQueryAndFind(queryParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, userQuests)
}

func CreateUserQuest(c *gin.Context) {
	userQuestRepo := infrastructure.NewUserQuestRepo(config.DB)
	userProgressReq := models.UserQuestProgressRequest{}

	if err := c.ShouldBindJSON(&userProgressReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var userQuest models.UserQuest
	if err := userQuestRepo.FindByUuidAndQid(&userProgressReq, &userQuest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, &userQuest)
}

func UpdateUserQuestProgress(c *gin.Context) {
	userQuestRepo := infrastructure.NewUserQuestRepo(config.DB)
	userRepo := infraUser.NewUserRepo(config.DB)
	userProgressReq := models.UserQuestProgressRequest{}

	if err := c.ShouldBindJSON(&userProgressReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var userQuest models.UserQuest
	if err := userQuestRepo.FindByUuidAndQid(&userProgressReq, &userQuest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := userQuestRepo.IncrementUserQuestProgression(&userQuest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user, err := userRepo.FindByUuid(userProgressReq.UserUuid.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if userQuest.IsCompleted == true {
		if err := userRepo.IncreaseUserXp(&userQuest, &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		if err := userQuestRepo.AddCurrencyAmountQuestToUser(&user, &userQuest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(200, &userQuest)
}
