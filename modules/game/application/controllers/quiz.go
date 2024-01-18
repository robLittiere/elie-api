package controllers

import (
	"elie-api/config"
	"elie-api/modules/game/infrastructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetQuizzes(c *gin.Context) {
	quizRepo := infrastructure.NewQuizRepo(config.DB)
	queryParams := c.Request.URL.Query()

	data, err := quizRepo.BuildQueryAndFind(queryParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, &data)
}
