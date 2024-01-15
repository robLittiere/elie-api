package controllers

import "github.com/gin-gonic/gin"

func GetQuests(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
