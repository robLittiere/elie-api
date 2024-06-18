package matchmaking

type QueueMessage struct {
	Type          string `json:"type"`
	Status        string `json:"status"`
	GameID        int    `json:"GameID"`
	QueuePosition int    `json:"queue_position"`
}

type RoomMessage struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	GameID int    `json:"GameID"`
	RoomID int    `json:"room_id"`
}

type DualQuizMessage struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	RoomID  int    `json:"room_id"`
	Message string `json:"message"`
}

type DualQuizGameMessage struct {
	Type     string `json:"type"`
	Status   string `json:"status"`
	RoomID   int    `json:"room_id"`
	QuizData string `json:"quiz_data"`
	Timer    int    `json:"timer"`
}
