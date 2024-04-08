package controllers

import (
	"elie-api/config"
	"elie-api/modules/gamification/infrastructure"
	"elie-api/modules/gamification/models"
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

func UpdateQuestProgress(c *gin.Context) {
	userQuestRepo := infrastructure.NewUserQuestRepo(config.DB)
	/*userRepo := infraUser.NewUserRepo(config.DB)*/
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

/*	user, err := userRepo.FindByUuid(userProgressReq.UserUuid.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}*/

	// TODO update user progression for a quest

	c.JSON(200, &userQuest)
}
