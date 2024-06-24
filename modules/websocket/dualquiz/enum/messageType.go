package enum

type MessageType int

const (
	DualQuizType MessageType = iota
	DualQuizAnswerType
	DualQuizEndType
)

func (m MessageType) String() string {
	switch m {
	case DualQuizType:
		return "dual_quiz"
	case DualQuizAnswerType:
		return "dual_quiz_answer"
	case DualQuizEndType:
		return "dual_quiz_end"
	}
	return "unknown_message_type"
}
