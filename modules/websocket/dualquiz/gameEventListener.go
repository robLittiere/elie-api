package dualquiz

type GameEventListener interface {
	OnPlayerCorrectAnswer(clientUuid string, roomId int)
	OnPlayerWrongAnswer(clientUuid string, roomId int)
	OnRoundEnd(roomId int)
	OnNextRoundStart(roomId int)
	OnGameEnd(roomId int, scoreMap MapPlayerData)
	OnRoomCleanup(roomId int)
}
