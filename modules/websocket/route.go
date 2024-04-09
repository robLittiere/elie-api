package websocket

import (
	"elie-api/modules/websocket/dualquiz"
	"elie-api/modules/websocket/matchmaking"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func WebsocketRoutes(router *gin.Engine, hubs map[string]Hub) {
	ws := router.Group("/ws")
	{
		ws.GET("/matchmaking/:game_id/:user_uuid", func(c *gin.Context) {
			gameID, err := strconv.Atoi(c.Param("game_id"))
			userUuid := c.Param("user_uuid")
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			matchmakingHub, ok := hubs["matchmaking"].(*matchmaking.Hub)
			if !ok {
				// handle error
				log.Printf("Error casting hub to matchmaking hub")
				c.JSON(500, gin.H{"error": "Errors occurred with matchmaking hub"})
			}
			matchmaking.ServeWsMatchmaking(matchmakingHub, c.Writer, c.Request, gameID, userUuid)
		})

		ws.GET("/dualquiz/:room_id/:user_uuid", func(c *gin.Context) {
			roomID, err := strconv.Atoi(c.Param("room_id"))
			userUuid := c.Param("user_uuid")
			if err != nil {
				fmt.Println("Room ID: ", roomID)
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			dqHub, ok := hubs["dualquiz"].(*dualquiz.DqHub)
			if !ok {
				// handle error
				log.Printf("Error casting hub to dualquiz hub")
				c.JSON(500, gin.H{"error": "Errors occurred with dualquiz hub"})
			}
			dualquiz.ServeWsDualQuiz(dqHub, c.Writer, c.Request, roomID, userUuid)
		})
	}
}
