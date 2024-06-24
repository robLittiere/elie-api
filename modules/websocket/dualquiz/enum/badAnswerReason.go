package enum

type BadAnswerReason int

const (
	IncorrectAnswer BadAnswerReason = iota
	Timeout
)

func (bar BadAnswerReason) String() string {
	switch bar {
	case IncorrectAnswer:
		return "IncorrectAnswer"
	case Timeout:
		return "Timeout"
	default:
		return "Unknown"
	}
}
