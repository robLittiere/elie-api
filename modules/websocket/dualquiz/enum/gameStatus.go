package enum

type GameStatus int

const (
	GameStarting GameStatus = iota
	GamePending
	GameFinished
)

func (gs GameStatus) String() string {
	switch gs {
	case GameStarting:
		return "Game is starting"
	case GamePending:
		return "Game is on pause"
	case GameFinished:
		return "Game is finished"
	}
	return "Unknown Status"
}
