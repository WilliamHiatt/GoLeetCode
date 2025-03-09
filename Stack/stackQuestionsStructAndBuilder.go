package Stack

type ValidParentheses struct {
	InString string
	Answer   bool
}

func ValidParenthesesBuilder(inString string, answer bool) *ValidParentheses {
	return &ValidParentheses{
		InString: inString,
		Answer:   answer,
	}
}
