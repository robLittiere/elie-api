package controllers

import "github.com/gin-gonic/gin"

func GetGames(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func CreateGame(c *gin.Context) {
	c.Status(201)
	return
}
