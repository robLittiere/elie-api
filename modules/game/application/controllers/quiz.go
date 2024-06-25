package controllers

import (
	"elie-api/config"
	"elie-api/modules/game/domain/service"
	"elie-api/modules/game/infrastructure"
	"elie-api/modules/game/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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

func CompleteUserQuiz(c *gin.Context) {
	var userquizRequest models.UserQuizRequest

	if err := c.ShouldBindJSON(&userquizRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userquiz = service.GetUserQuizFromRequest(config.DB, userquizRequest)

	if userquiz.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	//verify if the user has already completed the quiz
	userQuizService := service.NewUserQuizService(config.DB)
	exist, err := userQuizService.UserQuizExists(userquiz)
	fmt.Printf("exist: %v\n", exist)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if exist == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already completed this quiz"})
		return
	}

	quizRepo := infrastructure.NewQuizRepo(config.DB)
	err = quizRepo.CreateUserQuiz(&userquiz)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"New User Quiz created": userQuizService})
	return
}

func GetUserQuizzes(c *gin.Context) {
	userID := c.Param("id")

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	quizRepo := infrastructure.NewQuizRepo(config.DB)
	quizIds, err := quizRepo.FindQuizCompletedByUser(userIDInt)

	nextQuiz := quizRepo.FindNextQuiz(quizIds)
	log.Println("Next Quiz:", nextQuiz)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"quizIds":  quizIds,
		"nextQuiz": nextQuiz,
	})
}
