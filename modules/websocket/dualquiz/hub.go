package dualquiz

import (
	"elie-api/modules/websocket/matchmaking"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// DqHub maintains the set of active clients and broadcasts messages to the
// clients.
type DqHub struct {
	// Registered clients.
	clientsMux sync.RWMutex
	clients    map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Maps of rooms with players uuid inside
	roomsMux sync.RWMutex
	rooms    map[int][]string

	// Actual game rooms with players
	gameRoomsMux sync.RWMutex
	gameRooms    map[int][]*Client
}

func NewDqHub() *DqHub {
	return &DqHub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		rooms:      make(map[int][]string),
		gameRooms:  make(map[int][]*Client),
	}
}

func (dqh *DqHub) CreateRoom(client1 *matchmaking.Client, client2 *matchmaking.Client) int {
	// Create a room in the dualquiz hub
	dqh.roomsMux.Lock()
	defer dqh.roomsMux.Unlock()

	roomID := len(dqh.rooms)
	dqh.rooms[roomID] = []string{client1.UserUuid, client2.UserUuid}
	fmt.Printf("Room %d created with players %s and %s\n", roomID, client1.UserUuid, client2.UserUuid)
	return roomID

}

func (h *DqHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clientsMux.Lock()
			h.clients[client] = true
			client.CurrentStatus = StatusConnected
			h.clientsMux.Unlock()
			h.HandleConnection(client)
		case client := <-h.unregister:
			h.clientsMux.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			h.clientsMux.Unlock()
		case message := <-h.broadcast:
			h.clientsMux.RLock()
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.clientsMux.RUnlock()
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWsDualQuiz(hub *DqHub, w http.ResponseWriter, r *http.Request, roomId int, uuid string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	var client *Client
	client = &Client{hub: hub, conn: conn, send: make(chan []byte, 256), UserUuid: uuid, RoomID: roomId}

	client.hub.register <- client

	go client.writePump()
	go client.readPump()

}

func SendMessage(client *Client, msg interface{}) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}

	client.send <- jsonMsg
}

func (h *DqHub) isInARoom(client *Client) bool {
	h.gameRoomsMux.RLock()
	defer h.gameRoomsMux.RUnlock()
	for _, room := range h.gameRooms {
		for _, c := range room {
			if c.UserUuid == client.UserUuid {
				return true
			}
		}
	}
	return false
}

func (h *DqHub) isInRoom(client *Client, roomID int) bool {
	h.gameRoomsMux.RLock()
	defer h.gameRoomsMux.RUnlock()
	for _, c := range h.gameRooms[roomID] {
		if c.UserUuid == client.UserUuid {
			return true
		}
	}
	return false
}

func (h *DqHub) addClientToGameRoom(client *Client) error {
	h.gameRoomsMux.Lock()
	defer h.gameRoomsMux.Unlock()

	// First we check if the connection is legitimate
	if _, ok := h.rooms[client.RoomID]; !ok {
		return fmt.Errorf("room %d does not exist for user : %v", client.RoomID, client.UserUuid)
	}

	if _, ok := h.gameRooms[client.RoomID]; !ok {
		// Create game room if it does not exist yet
		h.gameRooms[client.RoomID] = []*Client{client}
	}

	// Check if user is already in the room and that there are less than 2 players in the room
	userInRoom := false
	for _, c := range h.gameRooms[client.RoomID] {
		if c.UserUuid == client.UserUuid {
			userInRoom = true
		}
	}
	if !userInRoom && len(h.gameRooms[client.RoomID]) < 2 {
		h.gameRooms[client.RoomID] = append(h.gameRooms[client.RoomID], client)
	}
	return nil
}

// TODO: Implement an observer to know the state of the rooms at each connection and disconnection
// TODO: Implement a function to launch the game when there are 2 ready players in the room
func (dqh *DqHub) HandleConnection(client *Client) {

	// Whenever we receive a connection, we need to attribute the client to its room
	// First we check if the connection is legitimate

	if err := dqh.addClientToGameRoom(client); err != nil {
		// Connexion is illegitimate
		fmt.Printf("Error adding client to game room\n %v\n", err)

		// Cut the connection
		connError := client.conn.Close()
		if connError != nil {
			log.Println(connError)
		}
	}

	if dqh.isInRoom(client, client.RoomID) {
		msg := matchmaking.DualQuizMessage{
			Type:   "DualQuiz",
			Status: "",
			RoomID: client.RoomID,
			Data:   "You have been put into a room",
		}
		SendMessage(client, msg)
		fmt.Printf("Client %s has joined room %d\n", client.UserUuid, client.RoomID)
		fmt.Printf("State of the room %d: %v users\n", client.RoomID, len(dqh.gameRooms[client.RoomID]))
	}

	// Check if there are now 2 users in the room...
	if len(dqh.gameRooms[client.RoomID]) == 2 {
		// We are ready to lauch the game
		msg := matchmaking.DualQuizMessage{
			Type:   "DualQuiz",
			Status: "Game",
			RoomID: client.RoomID,
			Data:   "Game is starting",
		}
		for _, c := range dqh.gameRooms[client.RoomID] {
			SendMessage(c, msg)
		}
		// Launch game
		// dqh.LaunchGame(client.RoomID)
	}
}
