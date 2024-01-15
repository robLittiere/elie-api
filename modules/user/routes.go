package user

import (
	"elie-api/cmd/api/modules/user/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	user := router.Group("/users")
	{
		user.GET("/", controllers.GetUsers)
		user.POST("/", controllers.CreateUser)
		user.POST("/signup")
		user.POST("/login")
	}
}
