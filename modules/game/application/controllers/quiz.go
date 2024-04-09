package controllers

import (
	"elie-api/config"
	"elie-api/modules/game/infrastructure"
	"elie-api/modules/game/models"
	infraUser "elie-api/modules/user/infrastructure"
	"fmt"
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

func CompletedQuizUser(c *gin.Context) {
	var userquiz models.UserQuiz
	var userquizRequest models.UserQuizRequest

	if err := c.ShouldBindJSON(&userquizRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("UserQuizRequest: %v\n", userquizRequest)

	// sortir le get user by uuid dans un service
	userRepo := infraUser.NewUserRepo(config.DB)
	user, err := userRepo.FindByUuid(userquizRequest.UserUuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userquiz.UserID = user.Id
	userquiz.QuizId = userquizRequest.QuizId

	// rajouter des cas d'excption dans le service exemple: si le quiz n'existe pas, si l'utilisateur n'existe pas, etc
	quizRepo := infrastructure.NewQuizRepo(config.DB)
	err = quizRepo.CreateUserQuiz(&userquiz)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
	return
}

func GetUserQuizzes(c *gin.Context) {
	userID := c.Param("id")
	quizRepo := infrastructure.NewQuizRepo(config.DB)
	quizzes, err := quizRepo.FindQuizzesCompletedByUser(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, &quizzes)
}
