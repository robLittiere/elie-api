package controllers

import (
	"elie-api/config"
	"elie-api/modules/game/infrastructure"
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

func CreateGame(c *gin.Context) {
	c.Status(201)
	return
}
