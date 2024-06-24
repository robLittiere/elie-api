package dualquiz

type GameEventListener interface {
	OnPlayerAnswer(clientUuid string, roomId int, isCorrect bool)
	OnRoundEnd(roomId int)
	OnNextRoundStart(roomId int)
	OnGameEnd(roomId int, scoreMap MapPlayerData)
	OnRoomCleanup(roomId int)
}
