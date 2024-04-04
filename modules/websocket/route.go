package websocket

import (
	"elie-api/modules/websocket/matchmaking"
	"github.com/gin-gonic/gin"
	"strconv"
)

func WebsocketRoutes(router *gin.Engine, hub *matchmaking.Hub) {
	ws := router.Group("/ws")
	{
		ws.GET("/matchmaking/:game_id/:user_uuid", func(c *gin.Context) {
			game_id, err := strconv.Atoi(c.Param("game_id"))
			user_uuid := c.Param("user_uuid")
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			matchmaking.ServeWsMatchmaking(hub, c.Writer, c.Request, game_id, user_uuid)
		})
	}
}
