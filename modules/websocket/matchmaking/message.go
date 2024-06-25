package matchmaking

type QueueMessage struct {
	Type          string      `json:"type"`
	Status        QueueStatus `json:"status"`
	StatusMessage string      `json:"status_message"`
	GameId        int         `json:"game_id"`
	QueuePosition int         `json:"queue_position"`
}

type RoomMessage struct {
	Type          string      `json:"type"`
	Status        QueueStatus `json:"status"`
	StatusMessage string      `json:"status_message"`
	GameId        int         `json:"game_id"`
	RoomId        int         `json:"room_id"`
	OpponentUuid  string      `json:"opponent_uuid"`
}
