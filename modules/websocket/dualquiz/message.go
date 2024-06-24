package dualquiz

import "elie-api/modules/websocket/dualquiz/enum"

type DualQuizMessage struct {
	Type          enum.MessageType `json:"type"`
	TypeMessage   string           `json:"type_message"`
	Status        enum.GameStatus  `json:"status"`
	StatusMessage string           `json:"status_message"`
	RoomID        int              `json:"room_id"`
	Message       string           `json:"message"`
}

type DualQuizGameMessage struct {
	Type            enum.MessageType `json:"type"`
	TypeMessage     string           `json:"type_message"`
	Status          enum.GameStatus  `json:"status"`
	StatusMessage   string           `json:"status_message"`
	RoomID          int              `json:"room_id"`
	QuizData        string           `json:"quiz_data"`
	CurrentQuestion int              `json:"current_question"`
	Timer           int              `json:"timer"`
}

type DualQuizGameAnswerMessage struct {
	Type           enum.MessageType `json:"type"`
	TypeMessage    string           `json:"type_message"`
	FromClientUuid string           `json:"client_uuid"`
	IsCorrect      bool             `json:"is_correct"`
	CorrectAnswer  int              `json:"correct_answer"`
	EndTimer       int              `json:"end_timer"`
}

type DualQuizGameEndMessage struct {
	Type          enum.MessageType `json:"type"`
	TypeMessage   string           `json:"type_message"`
	Status        enum.GameStatus  `json:"status"`
	StatusMessage string           `json:"status_message"`
	RoomId        int              `json:"room_id"`
	Winner        PlayerData       `json:"winner"`
	Loser         PlayerData       `json:"loser"`
	IsDraw        bool             `json:"is_equal"`
}

type ClientDualQuizMessage struct {
	Type   string `json:"type"`
	Choice int    `json:"choice"`
}
