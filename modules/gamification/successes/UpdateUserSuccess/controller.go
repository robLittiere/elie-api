package updateUserSuccess

import (
	"elie-api/modules/gamification/successes/domain/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateUserSuccessProgress(c *gin.Context) {
	userProgressReq := models.UserSuccessProgressRequest{}
	if err := c.ShouldBindJSON(&userProgressReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
}
