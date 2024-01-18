package user

import (
	"elie-api/modules/user/application/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	user := router.Group("/users")
	{
		user.GET("/", controllers.GetUsers)
		user.POST("/", controllers.CreateUser)
		user.POST("/signup", controllers.CreateUser)
		user.POST("/login")
	}
}
