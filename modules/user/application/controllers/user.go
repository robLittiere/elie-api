package controllers

import (
	"elie-api/config"
	"elie-api/modules/user/infrastructure"
	"elie-api/modules/user/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	userRepo := infrastructure.NewUserRepo(config.DB)
	queryParams := c.Request.URL.Query()

	users, err := userRepo.BuildQueryAndFind(queryParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Serialize users
	publicUsers := make([]models.PublicUser, 0)
	for _, user := range users {
		publicUsers = append(publicUsers, user.Serialize())
	}

	c.JSON(http.StatusOK, &publicUsers)
}

func CreateUser(c *gin.Context) {
	userRepo := infrastructure.NewUserRepo(config.DB)
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Printf("Missing request informations : %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	err := userRepo.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusCreated)
	return
}
