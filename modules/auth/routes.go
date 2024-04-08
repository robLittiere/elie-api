package auth

import (
	"elie-api/modules/auth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	{
		router.GET("/ping", controllers.PingHandler)
		router.POST("/signup", controllers.SignupHandler)
		router.POST("/login", controllers.LoginHandler)
	}

}
