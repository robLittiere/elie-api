package enum

type MessageType int

const (
	DualQuizType MessageType = iota
	DualQuizAnswerType
)

func (m MessageType) String() string {
	switch m {
	case DualQuizType:
		return "dual_quiz"
	case DualQuizAnswerType:
		return "dual_quiz_answer"
	}
	return "unknown_message_type"
}
