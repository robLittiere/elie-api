package tests

import (
	"elie-api/config"
	"elie-api/modules/auth"
	"elie-api/modules/game"
	"elie-api/modules/gamification/successes"
	"elie-api/modules/user"
	"elie-api/modules/websocket"
	"elie-api/modules/websocket/dualquiz"
	"elie-api/modules/websocket/matchmaking"
	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker/v2"
	"net/http"
	"net/http/httptest"
	"net/url"
)

var (
	c      *gin.Context
	w      *httptest.ResponseRecorder
	router *gin.Engine
	f      faker.Faker
)

func Init() {
	f = faker.New()
	config.SetUpTestDatabase()
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
	successes.GamificationRoutes(version)
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
