package controllers

import (
	"elie-api/config"
	"elie-api/modules/user/infrastructure"
	"elie-api/modules/user/models"
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

func GetUser(c *gin.Context) {
	userRepo := infrastructure.NewUserRepo(config.DB)
	uid := c.Param("uuid")

	user, err := userRepo.FindByUuid(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publicUser := user.Serialize()
	c.JSON(http.StatusOK, &publicUser)

}
