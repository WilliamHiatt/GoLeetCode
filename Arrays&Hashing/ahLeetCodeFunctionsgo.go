package arraysandhashing

func containsDuplicate(nums []int) bool {
	// https://leetcode.com/problems/contains-duplicate/description/
	alreadyPassed := make(map[int]bool)

	for _, val := range nums {
		if _, exists := alreadyPassed[val]; exists {
			return true
		} else {
			alreadyPassed[val] = true
		}
	}

	return false
}

func validAnagram(s string, t string) bool {
	// https://leetcode.com/problems/valid-anagram/description/
	if len(s) != len(t) {
		return false
	}

	firstStr := make(map[rune]int)

	for _, val := range s {
		firstStr[val]++
	}

	lettersNeeded := len(firstStr)
	for _, val := range t {
		if count, exists := firstStr[val]; exists && count > 0 {
			firstStr[val]--
			if firstStr[val] == 0 {
				lettersNeeded--
			}
		} else {
			return false
		}
	}

	return lettersNeeded == 0
}

func twoSum(nums []int, target int) []int {
	// https://leetcode.com/problems/two-sum/description/
	passed := make(map[int]int)

	for ind, val := range nums {
		wanted := target - val
		if value, exists := passed[wanted]; exists {
			return []int{value, ind}
		} else {
			passed[val] = ind
		}
	}

	return []int{0, 0}
}
