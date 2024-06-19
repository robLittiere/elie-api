package enum

type PlayerStatus int

const (
	PlayerStatusConnected PlayerStatus = iota
	PlayerStatusConnecting
	PlayerStatusDisconnected
	PlayerStatusReady
	PlayerStatusPlaying
)
