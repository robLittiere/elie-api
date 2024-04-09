package tests

import (
	"elie-api/modules/auth"
	"elie-api/modules/game"
	"elie-api/modules/gamification"
	"elie-api/modules/user"
	"elie-api/modules/websocket"
	"elie-api/modules/websocket/dualquiz"
	"elie-api/modules/websocket/matchmaking"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init Router
	router := InitRouter()

	router.Run(":8080")
}

func InitRouter() *gin.Engine {
	router := gin.Default()

	// Setup websocket matchHub
	matchHub := matchmaking.NewHub()
	dqHub := dualquiz.NewDqHub()
	matchHub.WithRoomCreator(dqHub)

	hubs := map[string]websocket.Hub{
		"matchmaking": matchHub,
		"dualquiz":    dqHub,
	}

	go matchHub.Run()
	go dqHub.Run()

	// Token test routes
	api := router.Group("/api")
	version := api.Group("/v1")

	// Group routes etc..
	auth.AuthRoutes(version)
	user.UserRoutes(version)
	gamification.GamificationRoutes(version)
	game.GameRoutes(version)
	websocket.WebsocketRoutes(router, hubs)

	return router
}

func CreateGinTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	// Use the router from InitRouter
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return c, w
}
