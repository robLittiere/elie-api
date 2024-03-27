package controllers

import (
	"elie-api/config"
	"elie-api/modules/gamification/infrastructure"
	"elie-api/modules/gamification/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
	if err := userSuccessRepo.FindByUuidAndQid(&userProgressReq, &userSuccess); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Launch update here
	// TODO update user progression for a quest

	c.JSON(200, &userSuccess)
}
