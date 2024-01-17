package controllers

import (
	"elie-api/config"
	"elie-api/modules/game/infrastructure"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetQuizzes(c *gin.Context) {
	quizRepo := infrastructure.NewQuizRepo(config.DB)

	data, err := quizRepo.Find()
	if err != nil {
		return
	}
	fmt.Println(data)
	c.JSON(200, &data)
}
