package dualquiz

import (
	"fmt"
	"sync"
)

type GameHandler struct {
	roomsMux sync.RWMutex
	rooms    map[int][]*GameClient
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		rooms: make(map[int][]*GameClient),
	}
}

func (gh *GameHandler) AddClientToRoom(roomId int, client *GameClient) {
	gh.roomsMux.Lock()
	defer gh.roomsMux.Unlock()

	if _, ok := gh.rooms[roomId]; !ok {
		gh.rooms[roomId] = make([]*GameClient, 0)
	}

	gh.rooms[roomId] = append(gh.rooms[roomId], client)
}

// LaunchGame launches the game for a given roomId
func (gh *GameHandler) LaunchGame(roomId int) {
	// Launch the game
}

// IsRoomReady checks if the room is ready for the game to be launched
func (gh *GameHandler) IsRoomReady(roomId int) bool {
	// Check if the room is ready, for now it's just 2 players, game is ready to be launched
	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	fmt.Printf("Checking if room %d is ready...\n", roomId)
	fmt.Printf("Room %d has %d players\n", roomId, len(gh.rooms[roomId]))
	return len(gh.rooms[roomId]) == 2
}

func (gh *GameHandler) IsPlayerWaitingForOpponent(id int, uuid string) bool {
	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	fmt.Printf("Checking if player %s is waiting for opponent in room %d\n", uuid, id)
	if len(gh.rooms[id]) == 1 {
		fmt.Printf("there is only one player in the game room indeed")
		return gh.rooms[id][0].UserUuid == uuid
	}
	return false
}
