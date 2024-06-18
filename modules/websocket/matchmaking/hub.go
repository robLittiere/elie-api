package matchmaking

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clientsMux sync.RWMutex
	clients    map[*Client]bool

	// Map websocket connection to client
	uuidToClientMux sync.RWMutex
	uuidToClient    map[string]*Client

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	queueHandler *QueueHandler

	roomCreator RoomCreator
}

func NewHub() *Hub {
	return &Hub{
		broadcast:    make(chan []byte),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		clients:      make(map[*Client]bool),
		uuidToClient: make(map[string]*Client),
		queueHandler: NewQueueHandler(),
	}
}

func (h *Hub) WithRoomCreator(roomCreator RoomCreator) {
	h.roomCreator = roomCreator
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clientsMux.Lock()
			h.clients[client] = true
			h.clientsMux.Unlock()
		case client := <-h.unregister:
			h.clientsMux.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				err := h.queueHandler.RemoveClientFromQueue(client.GameID, client.UserUuid)
				if err != nil {
					log.Printf("Error removing client from queue: %v", err)
				}
				h.RemoveClientFromUuidMap(client)
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
func ServeWsMatchmaking(hub *Hub, w http.ResponseWriter, r *http.Request, gameId int, uuid string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// TODO improve the queue system
	// TODO Send a message to the user once his position in the queue changes
	var client *Client
	// Check if user is already connected
	hub.uuidToClientMux.RLock()
	if _, ok := hub.uuidToClient[uuid]; ok {
		client = hub.uuidToClient[uuid]
	} else {
		client = &Client{hub: hub, conn: conn, send: make(chan []byte, 256), UserUuid: uuid, GameID: gameId}
		hub.uuidToClient[uuid] = client

		// Add client to the game queue and set the user's position in the queue
		hub.queueHandler.AddClientToQueueForGame(gameId, client)
		hub.queueHandler.CreateUserPositionInQueue(uuid, gameId)
	}
	hub.uuidToClientMux.RUnlock()

	client.hub.register <- client

	go client.writePump()
	go client.readPump()

	// Match new user to queued user if possible
	if err := hub.tryMatchClient(gameId); err != nil {
		log.Printf("Error trying to match client: %v", err)
	}

	// Check if client is in queue and has not been matched
	// We modify the queues when a match is made, so if he is still in the queue, he is still waiting
	if hub.queueHandler.IsClientInPositionQueue(client.UserUuid) {
		msg := QueueMessage{
			Type:          "Queue",
			Status:        WaitingInQueue,
			StatusMessage: WaitingInQueue.String(),
			GameId:        client.GameID,
		}
		SendMessage(client, msg)
	}
}

func (h *Hub) RemoveClientFromUuidMap(client *Client) {
	h.uuidToClientMux.Lock()
	defer h.uuidToClientMux.Unlock()

	if _, ok := h.uuidToClient[client.UserUuid]; ok {
		delete(h.uuidToClient, client.UserUuid)
	}
}

func (h *Hub) tryMatchClient(gameId int) error {
	// If there are at least 2 clients in the queue, try to match
	if h.queueHandler.CanMatchClientsForGame(gameId) {
		client1, client2 := h.queueHandler.GetFirstClientsInQueueForGame(gameId)
		err := h.queueHandler.HandleMatchClient(gameId, client1.UserUuid, client2.UserUuid)
		if err != nil {
			return err
		}

		// Create a private room
		// For now we can use a room creator, but we can also use a room factory or anything else to create rooms
		if h.roomCreator == nil {
			return fmt.Errorf("Room creator not set : Need to set a room creator to create a room")
		}
		roomId := h.roomCreator.CreateRoom(client1, client2)

		// Send a message to the clients containing the room id
		msg := RoomMessage{
			Type:          "Match",
			Status:        MatchedInQueue,
			StatusMessage: MatchedInQueue.String(),
			GameId:        gameId,
			RoomId:        roomId,
		}
		SendMessage(client1, msg)
		SendMessage(client2, msg)
	}
	return nil
}

func SendMessage(client *Client, msg interface{}) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}

	client.send <- jsonMsg
}
