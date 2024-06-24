package dualquiz

type GameClient struct {
	UserUuid             string
	RoomId               int
	AnswerTime           int
	HasAnsweredThisRound bool
	Score                int
}

// PlayerData is a struct that contains the necessary data to send to the clients.
// This is the minimum required to hold info about the player in the game
type PlayerData struct {
	UserUuid string
	Score    int
	IsWinner bool
}

// MapPlayerData is a struct that contains the winner and loser of the game
type MapPlayerData struct {
	Winner PlayerData
	Loser  PlayerData
}
