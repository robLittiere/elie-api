package controllers

import (
	"elie-api/config"
	"elie-api/modules/user/application/filters"
	"elie-api/modules/user/infrastructure"
	"elie-api/modules/user/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {

	userRepo := infrastructure.UserRepo{DB: config.DB}

	// get query params
	queryParams := c.Request.URL.Query()
	fmt.Println(queryParams)

	// For each queryParams, apply the criteria to our repo
	if len(queryParams) != 0 {
		for key, param := range queryParams {
			for _, value := range param {
				criteria, err := filters.GetUserFilter(key)
				fmt.Println(key, value)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				userRepo.ApplyCriteria(criteria, value)
			}
		}
	}

	// criteria, err := filters.GetUserFilter("username")

	// Now return users

	users, err := userRepo.Find()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("users found : %v", users)

	// Serialiaze users
	var publicUsers []models.PublicUser
	for i := 0; i < len(users); i++ {
		publicUsers = append(publicUsers, users[i].Serialize())
	}

	c.JSON(http.StatusOK, &publicUsers)
}

func CreateUser(c *gin.Context) {
	c.Status(http.StatusCreated)
	return
}
