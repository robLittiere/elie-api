package dualquiz

import (
	"elie-api/config"
	"elie-api/modules/game/infrastructure"
	"elie-api/modules/game/models"
	"fmt"
	"sync"
)

type GameHandler struct {
	roomsMux sync.RWMutex
	rooms    map[int]*Room
	quizRepo *infrastructure.QuizRepo
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		rooms:    make(map[int]*Room),
		quizRepo: infrastructure.NewQuizRepo(config.DB),
	}
}

type Room struct {
	Id      int
	Players []*GameClient
	Quiz    *models.Quiz
	Timer   int
}

func NewRoom(id int, players []*GameClient) *Room {
	return &Room{
		Id:      id,
		Players: players,
	}
}

func (gh *GameHandler) AddClientToRoom(roomId int, client *GameClient) {
	gh.roomsMux.Lock()
	defer gh.roomsMux.Unlock()

	if _, ok := gh.rooms[roomId]; !ok {
		// Create an empty room
		gh.rooms[roomId] = NewRoom(roomId, []*GameClient{})
	}

	gh.rooms[roomId].Players = append(gh.rooms[roomId].Players, client)

}

// SetQuizForRoom launches the game for a given roomId
func (gh *GameHandler) SetQuizForRoom(roomId int) {
	// Launch the game
	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	fmt.Printf("Launching game for room %d\n", roomId)
	// Get a random quiz
	quiz, err := gh.quizRepo.GetRandomQuiz()
	if err != nil {
		fmt.Printf("Error getting random quiz: %v\n", err)
		return
	}
	// Set the quiz to the room
	gh.rooms[roomId].Quiz = &quiz

}

// IsRoomReady checks if the room is ready for the game to be launched
func (gh *GameHandler) IsRoomReady(roomId int) bool {
	// Check if the room is ready, for now it's just 2 players, game is ready to be launched
	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	fmt.Printf("Checking if room %d is ready...\n", roomId)
	fmt.Printf("Room %d has %d players\n", roomId, len(gh.rooms[roomId].Players))
	return len(gh.rooms[roomId].Players) == 2
}

func (gh *GameHandler) IsPlayerWaitingForOpponent(id int, uuid string) bool {
	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	if len(gh.rooms[id].Players) == 1 {
		return gh.rooms[id].Players[0].UserUuid != uuid
	}
	return false
}

func (gh *GameHandler) GetQuizData(roomId int) *models.Quiz {

	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	return gh.rooms[roomId].Quiz

}
