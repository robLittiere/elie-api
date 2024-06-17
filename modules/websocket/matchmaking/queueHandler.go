package matchmaking

import (
	"fmt"
	"sync"
)

type QueueHandler struct {
	queueMux sync.RWMutex
	queue    map[int][]*Client

	userPosInQueueMux sync.RWMutex
	userPosInQueue    map[string]int
}

func (qHandler *QueueHandler) AddClientToQueueForGame(gameId int, client *Client) {
	qHandler.queueMux.Lock()
	defer qHandler.queueMux.Unlock()

	fmt.Printf("Adding client %v to queue for game %v\n", client.UserUuid, gameId)
	qHandler.queue[gameId] = append(qHandler.queue[gameId], client)
}

func (qHandler *QueueHandler) CreateUserPositionInQueue(uuid string, gameId int) {
	qHandler.userPosInQueueMux.Lock()
	defer qHandler.userPosInQueueMux.Unlock()

	fmt.Printf("Creating user position in queue for client %v\n", uuid)
	qHandler.userPosInQueue[uuid] = len(qHandler.queue[gameId]) - 1
}

func (qHandler *QueueHandler) RemoveClientFromQueue(clientGameId int, clientUuid string) error {
	qHandler.userPosInQueueMux.Lock()
	defer qHandler.userPosInQueueMux.Unlock()

	if _, ok := qHandler.userPosInQueue[clientUuid]; ok {
		gameId := clientGameId

		// Update queue games list
		qHandler.queue[gameId] = append(qHandler.queue[gameId][:qHandler.userPosInQueue[clientUuid]], qHandler.queue[gameId][qHandler.userPosInQueue[clientUuid]+1:]...)

		// Update user position in queue
		for i, client := range qHandler.queue[gameId] {
			qHandler.userPosInQueue[client.UserUuid] = i
		}
		return nil
	}
	return fmt.Errorf("client %v was not in queue, there might a problem with the queue system", clientUuid)
}

func (qHandler *QueueHandler) HandleMatchClient(gameId int, uuid1 string, uuid2 string) error {

	fmt.Printf("Attempting to match clients %v and %v for game %v...\n", uuid1, uuid2, gameId)

	qHandler.deleteClientInQueueFromUuid(gameId, uuid1)
	qHandler.deleteClientInQueueFromUuid(gameId, uuid2)

	// update clients position in queue
	qHandler.updateClientPositionsInQueue(gameId)

	// Remove clients from userPosInQueue
	qHandler.deleteClientInUserPosition(uuid1)
	qHandler.deleteClientInUserPosition(uuid2)

	return nil
}

func (qHandler *QueueHandler) CanMatchClientsForGame(gameId int) bool {
	qHandler.queueMux.RLock()
	defer qHandler.queueMux.RUnlock()
	return len(qHandler.queue[gameId]) >= 2
}

func (qHandler *QueueHandler) GetFirstClientsInQueueForGame(gameId int) (*Client, *Client) {
	qHandler.queueMux.RLock()
	defer qHandler.queueMux.RUnlock()
	client1 := qHandler.queue[gameId][0]
	client2 := qHandler.queue[gameId][1]
	return client1, client2
}

func NewQueueHandler() *QueueHandler {
	return &QueueHandler{
		queue:          make(map[int][]*Client),
		userPosInQueue: make(map[string]int),
	}
}

func (qHandler *QueueHandler) deleteClientInQueueFromUuid(gameId int, uuid string) {
	qHandler.queueMux.Lock()
	defer qHandler.queueMux.Unlock()

	fmt.Printf("Deleting client %v from queue for game %v...\n", uuid, gameId)

	for i, client := range qHandler.queue[gameId] {
		if client.UserUuid == uuid {
			qHandler.queue[gameId] = append(qHandler.queue[gameId][:i], qHandler.queue[gameId][i+1:]...)
			break
		}
	}
}

func (qHandler *QueueHandler) removeMatchedClientsFromQueue(gameId int) {
	qHandler.queueMux.Lock()
	defer qHandler.queueMux.Unlock()

	fmt.Printf("Removing matched clients from queue for game %v\n", gameId)

	qHandler.queue[gameId] = qHandler.queue[gameId][2:]

	for i, client := range qHandler.queue[gameId] {
		qHandler.userPosInQueue[client.UserUuid] = i
	}
}

func (qHandler *QueueHandler) updateClientPositionsInQueue(gameId int) {
	qHandler.queueMux.Lock()
	qHandler.userPosInQueueMux.Lock()
	defer qHandler.queueMux.Unlock()
	defer qHandler.userPosInQueueMux.Unlock()

	fmt.Printf("Updating client positions in queue for game %v...\n", gameId)

	for i, client := range qHandler.queue[gameId] {
		qHandler.userPosInQueue[client.UserUuid] = i
	}
}

func (qHandler *QueueHandler) deleteClientInUserPosition(uuid string) {
	qHandler.userPosInQueueMux.Lock()
	defer qHandler.userPosInQueueMux.Unlock()

	fmt.Printf("Deleting client %v from user position in queue\n", uuid)

	delete(qHandler.userPosInQueue, uuid)
}

func (qHandler *QueueHandler) IsClientInPositionQueue(uuid string) bool {
	qHandler.userPosInQueueMux.RLock()
	defer qHandler.userPosInQueueMux.RUnlock()

	if _, ok := qHandler.userPosInQueue[uuid]; !ok {
		return false
	}
	return true
}
