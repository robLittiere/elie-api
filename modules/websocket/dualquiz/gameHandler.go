package dualquiz

import (
	"elie-api/config"
	"elie-api/modules/game/infrastructure"
	"elie-api/modules/game/models"
	"elie-api/modules/websocket/dualquiz/enum"
	"fmt"
	"sync"
	"time"
)

type GameHandler struct {
	roomsMux       sync.RWMutex
	rooms          map[int]*Room
	quizRepo       *infrastructure.QuizRepo
	gEventListener GameEventListener
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

func NewGameHandler(gEventListener GameEventListener) *GameHandler {
	return &GameHandler{
		rooms:          make(map[int]*Room),
		quizRepo:       infrastructure.NewQuizRepo(config.DB),
		gEventListener: gEventListener,
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

func (gh *GameHandler) setHasAnsweredThisRound(roomId int, clientUuid string) {
	gh.roomsMux.Lock()
	defer gh.roomsMux.Unlock()

	for _, player := range gh.rooms[roomId].Players {
		if player.UserUuid == clientUuid {
			player.HasAnsweredThisRound = true
			break
		}
	}
}

func (gh *GameHandler) resetHasAnsweredThisRound(roomId int) {
	gh.roomsMux.Lock()
	defer gh.roomsMux.Unlock()

	for _, player := range gh.rooms[roomId].Players {
		player.HasAnsweredThisRound = false
	}
}

func (gh *GameHandler) setNextCurrentQuestion(roomId int) {
	gh.roomsMux.Lock()
	defer gh.roomsMux.Unlock()

	gh.rooms[roomId].CurrentQuestion++
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
func (gh *GameHandler) isAnswerCorrect(roomId int, answerIndex int) bool {

	// Get the current question
	return gh.GetCurrentCorrectAnswer(roomId) == answerIndex
}

func (gh *GameHandler) addToPlayerScore(roomId int, clientUuid string) {
	gh.roomsMux.Lock()
	defer gh.roomsMux.Unlock()

	for _, player := range gh.rooms[roomId].Players {
		if player.UserUuid == clientUuid {
			player.Score++
			break
		}
	}
}

func (gh *GameHandler) onPlayerAnswer(roomId int, clientUuid string, msg ClientDualQuizMessage) {
	isClientCorrect := gh.isAnswerCorrect(roomId, msg.Choice)
	fmt.Printf("Client %s answered %v\n", clientUuid, isClientCorrect)
	gh.setHasAnsweredThisRound(roomId, clientUuid)

	if isClientCorrect {
		gh.addToPlayerScore(roomId, clientUuid)
		gh.gEventListener.OnPlayerCorrectAnswer(clientUuid, roomId)
	} else {
		// Send a you were wrong message to the client with the correct answer in it
		gh.gEventListener.OnPlayerWrongAnswer(clientUuid, roomId)
	}

	// We can check if the round is over and end it
	if gh.isRoundOver(roomId) {
		gh.endRound(roomId)

		// If the game is over, we can end it
		if gh.isGameFinished(roomId) {
			gh.endGame(roomId)
		} else {
			gh.startNextRound(roomId)
		}
	}

}

// isRoundOver checks if all players have answered the current question
func (gh *GameHandler) isRoundOver(roomId int) bool {
	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	for _, player := range gh.rooms[roomId].Players {
		if !player.HasAnsweredThisRound {
			return false
		}
	}
	return true

}

func (gh *GameHandler) endRound(roomId int) {
	gh.gEventListener.OnRoundEnd(roomId)
}

// startNextRound starts the next round of the game, selects next question and reset all values
func (gh *GameHandler) startNextRound(roomId int) {
	gh.setNextCurrentQuestion(roomId)
	gh.updateCurrentCorrectAnswer(roomId)
	gh.resetHasAnsweredThisRound(roomId)

	go func() {
		time.Sleep(3 * time.Second)
		gh.gEventListener.OnNextRoundStart(roomId)
	}()
}

func (gh *GameHandler) isGameFinished(roomId int) bool {
	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	return gh.rooms[roomId].CurrentQuestion == len(gh.rooms[roomId].Quiz.Questions)-1
}

func (gh *GameHandler) endGame(roomId int) {
	mapPlayerData := gh.getFinalScoresMap(roomId)
	go func() {
		time.Sleep(3 * time.Second)
		gh.gEventListener.OnGameEnd(roomId, mapPlayerData)
	}()
}

func (gh *GameHandler) getFinalScoresMap(roomId int) MapPlayerData {
	var winner *GameClient
	var looser *GameClient

	gh.roomsMux.RLock()
	defer gh.roomsMux.RUnlock()

	for _, player := range gh.rooms[roomId].Players {
		if winner == nil || player.Score > winner.Score {
			winner = player
		}
		if looser == nil || player.Score < looser.Score {
			looser = player
		}
	}

	if winner == nil || looser == nil {
		// return an error or a default value
		fmt.Printf("Error getting winner and looser\n")
	}

	return MapPlayerData{
		Winner: PlayerData{
			UserUuid: winner.UserUuid,
			Score:    winner.Score,
			IsWinner: true,
		},
		Loser: PlayerData{
			UserUuid: looser.UserUuid,
			Score:    looser.Score,
			IsWinner: false,
		},
	}
}
