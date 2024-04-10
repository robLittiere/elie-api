package controllers

import (
	"elie-api/config"
	"elie-api/modules/gamification/infrastructure"
	"github.com/gin-gonic/gin"
)

func GetSuccess(c *gin.Context) {
	successRepo := infrastructure.NewSuccessRepo(config.DB)
	successes, err := successRepo.Find()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, successes)
}
