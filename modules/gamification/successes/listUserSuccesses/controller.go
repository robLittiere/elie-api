package listUserSuccesses

import (
	"elie-api/config"
	"elie-api/modules/gamification/successes/infra"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserSuccesses(c *gin.Context) {
	userSuccessRepo := infra.NewUserSuccessRepo(config.DB)

	queryParams := c.Request.URL.Query()

	userSuccesses, err := userSuccessRepo.BuildQueryAndFind(queryParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, userSuccesses)
}
