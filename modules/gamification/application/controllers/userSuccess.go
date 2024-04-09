package controllers

import (
	"elie-api/config"
	"elie-api/modules/gamification/infrastructure"
	"elie-api/modules/gamification/models"
	infrastructureUser "elie-api/modules/user/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

func GetUserSuccess(c *gin.Context) {
	userSuccessRepo := infrastructure.NewUserSuccessRepo(config.DB)

	queryParams := c.Request.URL.Query()

	userSuccess, err := userSuccessRepo.BuildQueryAndFind(queryParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, userSuccess)
}

func UpdateSuccessProgress(c *gin.Context) {
	userSuccessRepo := infrastructure.NewUserSuccessRepo(config.DB)
	userProgressReq := models.UserSuccessProgressRequest{}

	if err := c.ShouldBindJSON(&userProgressReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var userSuccess models.UserSuccess
	if err := userSuccessRepo.FindByUuidAndUserSuccessId(&userProgressReq, &userSuccess); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Launch update here
	// TODO update user progression for a quest

	c.JSON(200, &userSuccess)
}

func AddCurrencyAmountSuccessToUser(c *gin.Context) {
	successId := c.Param("success_id")
	uid := c.Param("uuid")

	userSuccessId, err := strconv.Atoi(successId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	userSuccessRepo := infrastructure.NewUserSuccessRepo(config.DB)
	userRepo := infrastructureUser.NewUserRepo(config.DB)
	userProgressReq := models.UserSuccessProgressRequest{UserUuid: uuid.MustParse(uid), UserSuccessId: userSuccessId}

	var userSuccess models.UserSuccess
	if err := userSuccessRepo.FindByUuidAndUserSuccessId(&userProgressReq, &userSuccess); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Check if the success is completed
	if userSuccess.IsCompleted {
		// Get the user details
		user, err := userRepo.FindByUuid(uid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Add CurrencyReward to CurrencyAmount
		user.CurrencyAmount += userSuccess.Success.CurrencyReward

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
