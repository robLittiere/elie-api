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

	queueMux sync.RWMutex
	queue    map[int][]*Client

	userPosInQueueMux sync.RWMutex
	userPosInQueue    map[string]int
}

func NewHub() *Hub {
	return &Hub{
		broadcast:      make(chan []byte),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		clients:        make(map[*Client]bool),
		uuidToClient:   make(map[string]*Client),
		queue:          make(map[int][]*Client),
		userPosInQueue: make(map[string]int),
	}
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
				err := h.RemoveClientFromQueue(client)
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

func (h *Hub) RemoveClientFromQueue(client *Client) error {
	h.userPosInQueueMux.Lock()
	defer h.userPosInQueueMux.Unlock()

	if _, ok := h.userPosInQueue[client.user_uuid]; ok {
		game_id := client.game_id

		// Update queue games list
		h.queue[game_id] = append(h.queue[game_id][:h.userPosInQueue[client.user_uuid]], h.queue[game_id][h.userPosInQueue[client.user_uuid]+1:]...)

		// Update user position in queue
		for i, client := range h.queue[game_id] {
			h.userPosInQueue[client.user_uuid] = i
		}
		return nil
	}
	return fmt.Errorf("client %v was not in queue, there might a problem with the queue system", client.user_uuid)

}

func (h *Hub) RemoveClientFromUuidMap(client *Client) {
	h.uuidToClientMux.Lock()
	defer h.uuidToClientMux.Unlock()

	if _, ok := h.uuidToClient[client.user_uuid]; ok {
		delete(h.uuidToClient, client.user_uuid)
	}
}

// serveWs handles websocket requests from the peer.
func ServeWsMatchmaking(hub *Hub, w http.ResponseWriter, r *http.Request, game_id int, uuid string) {
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
		client = &Client{hub: hub, conn: conn, send: make(chan []byte, 256), user_uuid: uuid, game_id: game_id}
		hub.uuidToClient[uuid] = client

		// Add client to the game queue and set the user's position in the queue
		hub.queueMux.Lock()
		hub.userPosInQueueMux.Lock()
		hub.queue[game_id] = append(hub.queue[game_id], client)
		hub.userPosInQueue[uuid] = len(hub.queue[game_id]) - 1
		hub.userPosInQueueMux.Unlock()
		hub.queueMux.Unlock()

	}
	hub.uuidToClientMux.RUnlock()

	client.hub.register <- client

	go client.writePump()
	go client.readPump()

	// Match new user to queued user if possible
	TryMatchClient(hub, game_id)

	// Check if client is in queue and has not been matched
	if _, ok := hub.userPosInQueue[uuid]; ok {
		msg := Message{
			Type:   "Queue",
			Status: "In Queue",
			GameID: game_id,
		}
		SendMessage(client, msg)
	}
}

func TryMatchClient(hub *Hub, game_id int) {
	hub.queueMux.Lock()
	hub.userPosInQueueMux.Lock()
	defer hub.userPosInQueueMux.Unlock()
	defer hub.queueMux.Unlock()

	if len(hub.queue[game_id]) >= 2 {
		client1 := hub.queue[game_id][0]
		client2 := hub.queue[game_id][1]
		hub.queue[game_id] = hub.queue[game_id][2:]

		// update clients position in queue
		for i, client := range hub.queue[game_id] {
			hub.userPosInQueue[client.user_uuid] = i
		}

		delete(hub.userPosInQueue, client1.user_uuid)
		delete(hub.userPosInQueue, client2.user_uuid)

		// Create and broadcast a message to the clients
		msg := Message{
			Type:   "Match",
			Status: "Match Found",
			GameID: game_id,
		}
		SendMessage(client1, msg)
		SendMessage(client2, msg)

	}
}

func SendMessage(client *Client, msg Message) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}

	client.send <- jsonMsg
}
