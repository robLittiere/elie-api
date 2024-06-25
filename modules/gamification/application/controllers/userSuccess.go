package controllers

import (
	"elie-api/config"
	"elie-api/modules/gamification/infrastructure"
	"elie-api/modules/gamification/models"
	infrastructureUser "elie-api/modules/user/infrastructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserSuccesses(c *gin.Context) {
	userSuccessRepo := infrastructure.NewUserSuccessRepo(config.DB)

	queryParams := c.Request.URL.Query()

	userSuccess, err := userSuccessRepo.BuildQueryAndFind(queryParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, userSuccess)
}

func UpdateUserSuccessProgress(c *gin.Context) {
	userSuccessRepo := infrastructure.NewUserSuccessRepo(config.DB)
	userProgressReq := models.UserSuccessProgressRequest{}
	userRepo := infrastructureUser.NewUserRepo(config.DB)

	if err := c.ShouldBindJSON(&userProgressReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var userSuccess models.UserSuccess
	if err := userSuccessRepo.FindByUuidAndUserSuccessId(&userProgressReq, &userSuccess); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := userSuccessRepo.IncrementUserSuccessProgression(&userSuccess); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := userRepo.FindByUuid(userProgressReq.UserUuid.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if userSuccess.IsCompleted == true {
		if err := userRepo.IncreaseUserXpSuccess(&userSuccess, &user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if err := userSuccessRepo.AddCurrencyAmountSuccessToUser(&user, &userSuccess); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if err := userSuccessRepo.AddNewUserSuccessByUser(&user, &userSuccess); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(200, &userSuccess)
}
