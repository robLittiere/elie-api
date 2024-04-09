package matchmaking

type RoomCreator interface {
	// CreateRoom creates a room with two clients and return the room id
	CreateRoom(client1 *Client, client2 *Client) int
}
