package dualquiz

import (
	"elie-api/modules/websocket/dualquiz/enum"
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
	// Registered clients and know if they're connected.
	clientsMux sync.RWMutex
	clients    map[*Client]bool

	// Map of clients by UUID.
	clientsByUuid map[string]*Client

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
	hub := &DqHub{
		broadcast:     make(chan []byte),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
		clientsByUuid: make(map[string]*Client),
		waitingRooms:  make(map[int][]string),
		gameRooms:     make(map[int][]*Client),
	}

	hub.gameHandler = NewGameHandler(hub)
	return hub
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

func (dqh *DqHub) getClientByUuid(uuid string) (*Client, bool) {
	dqh.clientsMux.RLock()
	defer dqh.clientsMux.RUnlock()

	client, ok := dqh.clientsByUuid[uuid]
	return client, ok
}

func (h *DqHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clientsMux.Lock()
			h.clients[client] = true
			h.clientsByUuid[client.UserUuid] = client
			h.clientsMux.Unlock()
		case client := <-h.unregister:
			h.clientsMux.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				delete(h.clientsByUuid, client.UserUuid)
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
	fmt.Printf("New client connected with user uuid %s\n", client.UserUuid)

	client.hub.register <- client

	// We need to save our instance of client in the hub
	h.addClientToGameRoom(client)

	err = h.HandleConnection(client)
	if err != nil {
		log.Println(err)
	}

	go client.writePump()
	go client.readPump()

}

// TODO: Implement an observer to know the state of the waitingRooms at each connection and disconnection
// TODO: Implement a function to launch the game when there are 2 ready players in the room
// Whenever we receive a connection, we need to attribute the client to its room
func (dqh *DqHub) HandleConnection(client *Client) error {

	var msg DualQuizMessage

	// First we check if the connection is legitimate
	if _, ok := dqh.waitingRooms[client.RoomID]; !ok {
		// Cut the connection
		connError := client.conn.Close()
		if connError != nil {
			log.Println(connError)
		}
		return fmt.Errorf("room %d does not exist for user : %v", client.RoomID, client.UserUuid)
	}

	dqh.gameHandler.AddClientToRoom(client.RoomID, &GameClient{
		UserUuid:             client.UserUuid,
		Score:                0,
		HasAnsweredThisRound: false,
	})

	fmt.Printf("Client %s has joined room %d\n", client.UserUuid, client.RoomID)

	// TODO Refactor game preparation, it is shiiit like that
	// If both players are in the room, game is ready to start
	if dqh.gameHandler.IsRoomReady(client.RoomID) {
		// Set the quiz for this room
		dqh.gameHandler.SetQuizForRoom(client.RoomID)
		// Launch game for this room
		dqh.LaunchGame(client.RoomID)
	}

	if dqh.gameHandler.IsPlayerWaitingForOpponent(client.RoomID, client.UserUuid) {
		roomStatus := dqh.gameHandler.GetRoomStatus(client.RoomID)
		msg = DualQuizMessage{
			Type:          enum.DualQuizType,
			TypeMessage:   enum.DualQuizType.String(),
			Status:        roomStatus,
			StatusMessage: roomStatus.String(),
			RoomID:        client.RoomID,
			Message:       "Waiting for opponent",
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

func (dqh *DqHub) LaunchGame(roomId int) {
	quizFromRoom := dqh.gameHandler.GetQuizData(roomId)
	dqh.gameHandler.SetRoomStatus(roomId, enum.GameStarting)

	msg := DualQuizGameMessage{
		Type:            enum.DualQuizType,
		TypeMessage:     enum.DualQuizType.String(),
		Status:          enum.GameStarting,
		StatusMessage:   enum.GameStarting.String(),
		RoomID:          roomId,
		QuizData:        quizFromRoom.ToJSON(),
		CurrentQuestion: dqh.gameHandler.GetCurrentQuestion(roomId),
		Timer:           0,
	}
	dqh.sendMessageToRoom(roomId, msg)
}

func (dqh *DqHub) HandleClientAnswer(c *Client, msgFromClient ClientDualQuizMessage) {
	dqh.gameHandler.onPlayerAnswer(c.RoomID, c.UserUuid, msgFromClient)
}

func (dqh *DqHub) OnPlayerCorrectAnswer(clientUuid string, roomId int) {
	// We need to send a message to the client that he was correct
	msg := DualQuizGameAnswerMessage{
		Type:           enum.DualQuizAnswerType,
		TypeMessage:    enum.DualQuizAnswerType.String(),
		FromClientUuid: clientUuid,
		IsCorrect:      true,
		CorrectAnswer:  dqh.gameHandler.GetCurrentCorrectAnswer(roomId),
		EndTimer:       0,
	}

	dqh.sendMessageToRoom(roomId, msg)
}

func (dqh *DqHub) OnPlayerWrongAnswer(clientUuid string, roomId int) {
	// We need to send a message to the client that he was wrong
	msg := DualQuizGameAnswerMessage{
		Type:           enum.DualQuizAnswerType,
		TypeMessage:    enum.DualQuizAnswerType.String(),
		FromClientUuid: clientUuid,
		IsCorrect:      false,
		CorrectAnswer:  dqh.gameHandler.GetCurrentCorrectAnswer(roomId),
		EndTimer:       0,
	}

	dqh.sendMessageToRoom(roomId, msg)
}

func (dqh *DqHub) OnRoundEnd(roomId int) {
	// We need to send a message to the clients that the round has ended
	msg := DualQuizMessage{
		Type:          enum.DualQuizType,
		TypeMessage:   enum.DualQuizType.String(),
		Status:        enum.GameRoundFinished,
		StatusMessage: enum.GameRoundFinished.String(),
		RoomID:        roomId,
		Message:       "Round has ended",
	}

	dqh.sendMessageToRoom(roomId, msg)

}

func (dqh *DqHub) OnNextRoundStart(roomId int) {
	// We need to send a message to the clients that the next round has started
	msg := DualQuizMessage{
		Type:          enum.DualQuizType,
		TypeMessage:   enum.DualQuizType.String(),
		Status:        enum.GameRoundStarting,
		StatusMessage: enum.GameRoundStarting.String(),
		RoomID:        roomId,
		Message:       "Next round has started",
	}

	dqh.sendMessageToRoom(roomId, msg)

	quizFromRoom := dqh.gameHandler.GetQuizData(roomId)
	message := DualQuizGameMessage{
		Type:            enum.DualQuizType,
		TypeMessage:     enum.DualQuizType.String(),
		Status:          enum.GameStarting,
		StatusMessage:   enum.GameStarting.String(),
		RoomID:          roomId,
		QuizData:        quizFromRoom.ToJSON(),
		CurrentQuestion: dqh.gameHandler.GetCurrentQuestion(roomId),
		Timer:           0,
	}
	dqh.sendMessageToRoom(roomId, message)

}
