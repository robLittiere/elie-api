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

	// Maps of waitingRooms with players uuid inside
	waitingRoomsMux sync.RWMutex
	waitingRooms    map[int][]string

	// Actual game waitingRooms with clients for this hub
	gameRoomsMux sync.RWMutex
	gameRooms    map[int][]*Client

	gameHandler *GameHandler
}

func NewDqHub() *DqHub {
	return &DqHub{
		broadcast:    make(chan []byte),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		clients:      make(map[*Client]bool),
		waitingRooms: make(map[int][]string),
		gameRooms:    make(map[int][]*Client),
		gameHandler:  NewGameHandler(),
	}
}

func (dqh *DqHub) CreateRoom(client1 *matchmaking.Client, client2 *matchmaking.Client) int {
	// Create a room in the dualquiz hub
	dqh.waitingRoomsMux.Lock()
	defer dqh.waitingRoomsMux.Unlock()

	roomID := len(dqh.waitingRooms)
	dqh.waitingRooms[roomID] = []string{client1.UserUuid, client2.UserUuid}
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
func ServeWsDualQuiz(h *DqHub, w http.ResponseWriter, r *http.Request, roomId int, uuid string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	var client *Client
	client = &Client{hub: h, conn: conn, send: make(chan []byte, 256), UserUuid: uuid, RoomID: roomId}

	client.hub.register <- client

	err = h.HandleConnection(client)
	if err != nil {
		log.Println(err)
	}
	// We need to save our instance of client in the hub
	h.addClientToGameRoom(client)

	go client.writePump()
	go client.readPump()

}

// TODO: Implement an observer to know the state of the waitingRooms at each connection and disconnection
// TODO: Implement a function to launch the game when there are 2 ready players in the room
// Whenever we receive a connection, we need to attribute the client to its room
func (dqh *DqHub) HandleConnection(client *Client) error {

	var msg matchmaking.DualQuizMessage

	// First we check if the connection is legitimate
	if _, ok := dqh.waitingRooms[client.RoomID]; !ok {
		// Cut the connection
		connError := client.conn.Close()
		if connError != nil {
			log.Println(connError)
		}
		return fmt.Errorf("room %d does not exist for user : %v", client.RoomID, client.UserUuid)
	}

	dqh.gameHandler.AddClientToRoom(client.RoomID, &GameClient{UserUuid: client.UserUuid})

	fmt.Printf("Client %s has joined room %d\n", client.UserUuid, client.RoomID)

	if dqh.gameHandler.IsRoomReady(client.RoomID) {
		// We are ready to launch the game

		msg = matchmaking.DualQuizMessage{
			Type:    "DualQuiz",
			Status:  "Game",
			RoomID:  client.RoomID,
			Message: "Game is starting",
		}
		dqh.sendMessageToRoom(client.RoomID, msg)

		// Launch game
		dqh.gameHandler.LaunchGame(client.RoomID)
	}

	if dqh.gameHandler.IsPlayerWaitingForOpponent(client.RoomID, client.UserUuid) {
		msg = matchmaking.DualQuizMessage{
			Type:    "DualQuiz",
			Status:  "Pending",
			RoomID:  client.RoomID,
			Message: "Waiting for opponent",
		}
		dqh.sendMessage(client, msg)
	}
	return nil
}

func (dqh *DqHub) sendMessageToRoom(roomId int, msg interface{}) {
	dqh.gameRoomsMux.RLock()
	defer dqh.gameRoomsMux.RUnlock()

	for _, client := range dqh.gameRooms[roomId] {
		dqh.sendMessage(client, msg)
	}
}

func (dqh *DqHub) sendMessage(client *Client, msg interface{}) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}

	client.send <- jsonMsg
}

func (dqh *DqHub) addClientToGameRoom(client *Client) {
	dqh.gameRoomsMux.Lock()
	defer dqh.gameRoomsMux.Unlock()

	if _, ok := dqh.gameRooms[client.RoomID]; !ok {
		dqh.gameRooms[client.RoomID] = make([]*Client, 0)
	}

	dqh.gameRooms[client.RoomID] = append(dqh.gameRooms[client.RoomID], client)

}
