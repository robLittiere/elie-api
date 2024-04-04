package controllers

import (
	"elie-api/config"
	"elie-api/modules/user/infrastructure"
	"elie-api/modules/user/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	userRepo := infrastructure.NewUserRepo(config.DB)
	var userRequest models.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		fmt.Printf("Missing request informations : %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		err := c.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			fmt.Printf("Error while aborting request : %v\n", err.Error())
		}
		return
	}

	user, err := userRepo.FindByEmail(userRequest.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	publicUser := user.Serialize()
	c.JSON(http.StatusOK, &publicUser)
}

func SignupHandler(c *gin.Context) {
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

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
