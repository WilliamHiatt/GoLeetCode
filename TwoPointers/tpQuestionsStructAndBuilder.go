package TwoPointers

type ValidPalindrome struct {
	InString string
	Answer   bool
}

func ValidPalindromeBuilder(inString string, answer bool) *ValidPalindrome {
	return &ValidPalindrome{
		InString: inString,
		Answer:   answer,
	}
}
