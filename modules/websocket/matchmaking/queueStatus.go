package matchmaking

type QueueStatus int

const (
	WaitingInQueue QueueStatus = iota
	MatchedInQueue
)

func (qs QueueStatus) String() string {
	switch qs {
	case WaitingInQueue:
		return "Waiting for a match"
	case MatchedInQueue:
		return "Matched Found"
	}
	return "Unknown Status"
}
