package Stack

func IsValidParentheses(s string) bool {
	stack := []rune{}
	key := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{'}

	for _, char := range s {
		if value, exists := key[char]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != value {
				return false
			}

			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
