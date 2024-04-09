package controllers

import (
	"elie-api/config"
	"elie-api/modules/gamification/infrastructure"
	"elie-api/modules/gamification/models"
	infraUser "elie-api/modules/user/infrastructure"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
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
	}

	c.JSON(200, &userQuest)
}

func AddCurrencyAmountQuestToUser(c *gin.Context) {
	questId := c.Param("quest_id")
	uid := c.Param("uuid")

	fmt.Println("questId: ", questId)
	fmt.Println("uid: ", uid)

	userQuestId, err := strconv.Atoi(questId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	userQuestRepo := infrastructure.NewUserQuestRepo(config.DB)
	userRepo := infraUser.NewUserRepo(config.DB)
	userProgressReq := models.UserQuestProgressRequest{UserUuid: uuid.MustParse(uid), UserQuestId: userQuestId}

	var userQuest models.UserQuest
	if err := userQuestRepo.FindByUuidAndQid(&userProgressReq, &userQuest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Check if the success is completed
	if userQuest.IsCompleted {
		// Get the user details
		user, err := userRepo.FindByUuid(uid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Add CurrencyReward to CurrencyAmount
		user.CurrencyAmount += userQuest.Quest.CurrencyReward

		// Update the user's CurrencyAmount
		if err := userRepo.UpdateUser(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user"})
			return
		}

		// Return success message
		c.JSON(http.StatusOK, gin.H{"message": "Currency added successfully"})
		return
	}

	//c.JSON(http.StatusBadRequest, gin.H{"message": "Success is not completed yet"})
}
