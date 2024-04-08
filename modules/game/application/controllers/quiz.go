package controllers

import (
	"elie-api/config"
	"elie-api/modules/game/infrastructure"
	"elie-api/modules/game/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetQuizGameData(c *gin.Context) {
	// Get the quiz game id from the url as well as the query parameters
	id := c.Param("id")
	queryParams := c.Request.URL.Query()

	quizRepo := infrastructure.NewQuizRepo(config.DB)

	data, err := quizRepo.BuildQueryAndFindByData(id, queryParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, &data)
}

func CreateUserQuizz(c *gin.Context) {
	userquizz := models.UserQuiz{}
	if err := c.ShouldBindJSON(&userquizz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	quizRepo := infrastructure.NewQuizRepo(config.DB)
	err := quizRepo.CreateQuizGame(&userquizz)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &userquizz)
}
