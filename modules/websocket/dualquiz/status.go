package dualquiz

type Status int

const (
	StatusConnected Status = iota
	StatusConnecting
	StatusDisconnected
	StatusReady
	StatusPlaying
)
