package dualquiz

type DualQuizMessage struct {
	Type          string     `json:"type"`
	Status        GameStatus `json:"status"`
	StatusMessage string     `json:"status_message"`
	RoomID        int        `json:"room_id"`
	Message       string     `json:"message"`
}

type DualQuizGameMessage struct {
	Type          string     `json:"type"`
	Status        GameStatus `json:"status"`
	StatusMessage string     `json:"status_message"`
	RoomID        int        `json:"room_id"`
	QuizData      string     `json:"quiz_data"`
	Timer         int        `json:"timer"`
}
