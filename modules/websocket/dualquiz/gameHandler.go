package dualquiz

import (
	"elie-api/config"
	"elie-api/modules/game/infrastructure"
	"elie-api/modules/game/models"
	"elie-api/modules/websocket/dualquiz/enum"
	"fmt"
	"sync"
)

type GameHandler struct {
	roomsMux sync.RWMutex
	rooms    map[int]*Room
	quizRepo *infrastructure.QuizRepo
}

type Room struct {
	Id                   int
	Players              []*GameClient
	Quiz                 *models.Quiz
	CurrentQuestion      int
	CurrentCorrectAnswer int
	Timer                int
	Status               enum.GameStatus
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		rooms:    make(map[int]*Room),
		quizRepo: infrastructure.NewQuizRepo(config.DB),
	}
}

func NewRoom(id int, players []*GameClient) *Room {
	return &Room{
		Id:      id,
		Players: players,
		Status:  enum.GamePending,
	}
}

func (gh *GameHandler) AddClientToRoom(roomId int, client *GameClient) {
	gh.roomsMux.Lock()
	defer gh.roomsMux.Unlock()

	if _, ok := gh.rooms[roomId]; !ok {
		// Create an empty room
		gh.rooms[roomId] = NewRoom(roomId, []*GameClient{})
	}

	gh.rooms[roomId].Players = append(gh.rooms[roomId].Players, client)

}

func (gh *GameHandler) GetRoomStatus(roomId int) enum.GameStatus {
	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	return gh.rooms[roomId].Status
}

func (gh *GameHandler) GetCurrentQuestion(roomId int) int {
	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	return gh.rooms[roomId].CurrentQuestion
}

func (gh *GameHandler) GetQuizData(roomId int) *models.Quiz {
	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	return gh.rooms[roomId].Quiz
}

func (gh *GameHandler) GetCurrentCorrectAnswer(roomId int) int {
	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	return gh.rooms[roomId].CurrentCorrectAnswer
}

func (gh *GameHandler) setQuizData(roomId int, quiz *models.Quiz) {
	gh.roomsMux.Lock()
	defer gh.roomsMux.Unlock()

	gh.rooms[roomId].Quiz = quiz
}

func (gh *GameHandler) setStartingStatus(roomId int) {
	gh.roomsMux.Lock()
	defer gh.roomsMux.Unlock()

	gh.rooms[roomId].Status = enum.GameStarting
}

// SetQuizForRoom launches the game for a given roomId
func (gh *GameHandler) SetQuizForRoom(roomId int) {
	// Launch the game
	fmt.Printf("Launching game for room %d\n", roomId)

	// Get a random quiz
	quiz, err := gh.quizRepo.GetRandomQuiz()
	if err != nil {
		fmt.Printf("Error getting random quiz: %v\n", err)
		return
	}

	// Set the quiz to the room
	gh.setQuizData(roomId, &quiz)
	gh.setCurrentQuestion(roomId, 0)
	gh.updateCurrentCorrectAnswer(roomId)
}

func (gh *GameHandler) setCurrentQuestion(roomId int, question int) {
	gh.roomsMux.Lock()
	defer gh.roomsMux.Unlock()

	gh.rooms[roomId].CurrentQuestion = question
}

func (gh *GameHandler) setCurrentCorrectAnswer(roomId int, currentAnswerIndex int) {
	gh.roomsMux.Lock()
	defer gh.roomsMux.Unlock()

	gh.rooms[roomId].CurrentCorrectAnswer = currentAnswerIndex
}

func (gh *GameHandler) updateCurrentCorrectAnswer(roomId int) {
	currentQuestion := gh.GetCurrentQuestion(roomId)
	quiz := gh.GetQuizData(roomId)

	answers := quiz.Questions[currentQuestion].Answers
	for i, answer := range answers {
		if quiz.Questions[currentQuestion].GoodAnswer == answer {
			gh.setCurrentCorrectAnswer(roomId, i)
			break
		}
	}
}

func (gh *GameHandler) SetRoomStatus(roomId int, status enum.GameStatus) {
	gh.roomsMux.Lock()
	defer gh.roomsMux.Unlock()

	gh.rooms[roomId].Status = status
}

// IsRoomReady checks if the room is ready for the game to be launched
func (gh *GameHandler) IsRoomReady(roomId int) bool {
	// Check if the room is ready, for now it's just 2 players, game is ready to be launched
	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	fmt.Printf("Checking if room %d is ready...\n", roomId)
	fmt.Printf("Room %d has %d players\n", roomId, len(gh.rooms[roomId].Players))
	return len(gh.rooms[roomId].Players) == 2
}

func (gh *GameHandler) IsPlayerWaitingForOpponent(id int, uuid string) bool {
	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	if len(gh.rooms[id].Players) == 1 {
		return gh.rooms[id].Players[0].UserUuid != uuid
	}
	return false
}

// IsAnswerCorrect Determine if the answer given to the current question is correct
func (gh *GameHandler) IsAnswerCorrect(roomId int, msg ClientDualQuizMessage) bool {

	// Get the current question
	return gh.GetCurrentCorrectAnswer(roomId) == msg.Choice
}

func (gh *GameHandler) AddToPlayerScore(roomId int, clientUuid string) {
	gh.roomsMux.Lock()
	defer gh.roomsMux.Unlock()

	for _, player := range gh.rooms[roomId].Players {
		if player.UserUuid == clientUuid {
			player.Score++
			break
		}
	}
}
