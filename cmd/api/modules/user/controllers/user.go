package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func CreateUser(c *gin.Context) {
	c.Status(http.StatusCreated)
	return
}
