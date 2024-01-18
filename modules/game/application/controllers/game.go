package controllers

import (
	"elie-api/config"
	"elie-api/modules/game/infrastructure"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGames(c *gin.Context) {
	gameRepo := infrastructure.NewGameRepo(config.DB)
	queryParams := c.Request.URL.Query()

	games, err := gameRepo.BuildQueryAndFind(queryParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &games)
}

func GetQuizGames(c *gin.Context) {

	gameRepo := infrastructure.NewGameRepo(config.DB)

	quizzes, err := gameRepo.FindQuizGames()
	if err != nil {
		return
	}
	fmt.Println(quizzes)
	c.JSON(200, &quizzes)
}

func CreateGame(c *gin.Context) {
	c.Status(201)
	return
}
