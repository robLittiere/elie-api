package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(router *gin.RouterGroup) {
	{
		router.GET("/ping")
		router.POST("/signup")
		router.POST("/login")
	}

}
