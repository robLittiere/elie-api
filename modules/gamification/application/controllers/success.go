package controllers

import (
	"elie-api/config"
	"elie-api/modules/gamification/infrastructure"
	"github.com/gin-gonic/gin"
)

func GetSuccesses(c *gin.Context) {
	successRepo := infrastructure.NewSuccessRepo(config.DB)
	queryParams := c.Request.URL.Query()

	successes, err := successRepo.BuildQueryAndFind(queryParams)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, successes)
}
