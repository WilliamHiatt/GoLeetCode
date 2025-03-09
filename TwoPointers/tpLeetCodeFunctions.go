package TwoPointers

func validPalindrome(s string) bool {
	// https://leetcode.com/problems/valid-palindrome/description/
	leftPtr := 0
	rightPtr := len(s) - 1

	for leftPtr < rightPtr {

		for leftPtr < rightPtr && !validPalindromeHelper(s[leftPtr]) {
			leftPtr++
		}
		for leftPtr < rightPtr && !validPalindromeHelper(s[rightPtr]) {
			rightPtr--
		}

		if toLower(s[leftPtr]) != toLower(s[rightPtr]) {
			return false
		}

		leftPtr++
		rightPtr--
	}

	return true
}

func validPalindromeHelper(incoming byte) bool {
	if ('a' <= incoming && 'z' >= incoming) || ('A' <= incoming && 'Z' >= incoming) || ('0' <= incoming && '9' >= incoming) {
		return true
	}

	return false
}

func toLower(incoming byte) byte {
	if incoming >= 'A' && incoming <= 'Z' {
		return incoming + ('a' - 'A')
	}

	return incoming
}
