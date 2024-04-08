package matchmaking

type Message struct {
	Type          string `json:"type"`
	Status        string `json:"status"`
	GameID        int    `json:"game_id"`
	QueuePosition int    `json:"queue_position"`
}
