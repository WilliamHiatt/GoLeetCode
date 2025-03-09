package arraysandhashing

import (
	"fmt"
	"testing"
)

func arraysandhashing() {
	fmt.Println("HELLO")
}

func TestContainsDuplicateQuestions(t *testing.T) {
	// https://leetcode.com/problems/contains-duplicate/description/
	noDuplicates := NewDuplicateQuestion([]int{1, 2, 3, 4, 5}, false)
	withDuplicates := NewDuplicateQuestion([]int{1, 2, 3, 4, 5, 3}, true)
	allDuplicates := NewDuplicateQuestion([]int{7, 7, 7, 7, 7}, true)
	largeArray := NewDuplicateQuestion([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 10}, true)
	singleElement := NewDuplicateQuestion([]int{42}, false)

	if returned := containsDuplicate(noDuplicates.ArrayOne); returned != noDuplicates.Answer {
		t.Errorf("noDuplicates incorrect result, got: %v, want:%v", returned, noDuplicates.Answer)
	}

	if returned := containsDuplicate(withDuplicates.ArrayOne); returned != withDuplicates.Answer {
		t.Errorf("withDuplicates incorrect result, got: %v, want:%v", returned, withDuplicates.Answer)
	}

	if returned := containsDuplicate(allDuplicates.ArrayOne); returned != allDuplicates.Answer {
		t.Errorf("allDuplicates incorrect result, got: %v, want:%v", returned, allDuplicates.Answer)
	}

	if returned := containsDuplicate(largeArray.ArrayOne); returned != largeArray.Answer {
		t.Errorf("largeArray incorrect result, got: %v, want:%v", returned, largeArray.Answer)
	}

	if returned := containsDuplicate(singleElement.ArrayOne); returned != singleElement.Answer {
		t.Errorf("singleElement incorrect result, got: %v, want:%v", returned, singleElement.Answer)
	}
}

func TestValidAnagramQuestions(t *testing.T) {
	// https://leetcode.com/problems/valid-ankagram/description/
	validAnagramOne := NewValidAnagram("listen", "silent", true)
	validAnagramTwo := NewValidAnagram("triangle", "integral", true)
	validAnagramThree := NewValidAnagram("conversation", "voicerantson", true)

	invalidAnagramOne := NewValidAnagram("hello", "world", false)
	invalidAnagramTwo := NewValidAnagram("rat", "car", false)
	invalidAnagramThree := NewValidAnagram("restup", "let", false)
	invalidAnagramFour := NewValidAnagram("yup", "never", false)

	if returned := validAnagram(validAnagramOne.FirstString, validAnagramOne.SecondString); returned != validAnagramOne.Answer {
		t.Errorf("ValidAnagramOne failed, got:%v, want:%v", returned, validAnagramOne.Answer)
	}

	if returned := validAnagram(validAnagramTwo.FirstString, validAnagramTwo.SecondString); returned != validAnagramTwo.Answer {
		t.Errorf("validAnagramTwo failed, got:%v, want:%v", returned, validAnagramTwo.Answer)
	}

	if returned := validAnagram(validAnagramThree.FirstString, validAnagramThree.SecondString); returned != validAnagramThree.Answer {
		t.Errorf("validAnagramThree failed, got:%v, want:%v", returned, validAnagramThree.Answer)
	}

	if returned := validAnagram(invalidAnagramOne.FirstString, invalidAnagramOne.SecondString); returned != invalidAnagramOne.Answer {
		t.Errorf("invalidAnagramOne failed, got:%v, want:%v", returned, invalidAnagramOne.Answer)
	}

	if returned := validAnagram(invalidAnagramTwo.FirstString, invalidAnagramTwo.SecondString); returned != invalidAnagramTwo.Answer {
		t.Errorf("invalidAnagramTwo failed, got:%v, want:%v", returned, invalidAnagramTwo.Answer)
	}

	if returned := validAnagram(invalidAnagramThree.FirstString, invalidAnagramThree.SecondString); returned != invalidAnagramThree.Answer {
		t.Errorf("invalidAnagramThree failed, got:%v, want:%v", returned, invalidAnagramThree.Answer)
	}

	if returned := validAnagram(invalidAnagramFour.FirstString, invalidAnagramFour.SecondString); returned != invalidAnagramFour.Answer {
		t.Errorf("invalidAnagramFour failed, got:%v, want:%v", returned, invalidAnagramFour.Answer)
	}
}

func TestTwoSumQuestions(t *testing.T) {
	// https://leetcode.com/problems/two-sum/description/
	twoSumOne := TwoSumQuestion([]int{12, 87, 4, 29, 65}, 41, []int{0, 3})
	twoSumTwo := TwoSumQuestion([]int{56, 91, 3}, 94, []int{1, 2})
	twoSumThree := TwoSumQuestion([]int{17, 49, 72, 8, 63, 54, 31}, 71, []int{3, 4})
	twoSumFour := TwoSumQuestion([]int{38, 51, 9, 27, 73, 45}, 72, []int{3, 5})
	twoSumFive := TwoSumQuestion([]int{99, 28, 11, 45, 68, 3, 72, 60}, 159, []int{0, 7})
	twoSumSix := TwoSumQuestion([]int{21, 54, 38, 87}, 92, []int{1, 2})

	if returned := twoSum(twoSumOne.Array, twoSumOne.Target); !twoSumOne.MatchesAnswer(returned) {
		t.Errorf("twoSumOne failed, got:%v, want:%v", returned, twoSumOne.Answer)
	}

	if returned := twoSum(twoSumTwo.Array, twoSumTwo.Target); !twoSumTwo.MatchesAnswer(returned) {
		t.Errorf("twoSumTwo failed, got:%v, want:%v", returned, twoSumTwo.Answer)
	}

	if returned := twoSum(twoSumThree.Array, twoSumThree.Target); !twoSumThree.MatchesAnswer(returned) {
		t.Errorf("twoSumThree failed, got:%v, want:%v", returned, twoSumThree.Answer)
	}

	if returned := twoSum(twoSumFour.Array, twoSumFour.Target); !twoSumFour.MatchesAnswer(returned) {
		t.Errorf("twoSumFour failed, got:%v, want:%v", returned, twoSumFour.Answer)
	}

	if returned := twoSum(twoSumFive.Array, twoSumFive.Target); !twoSumFive.MatchesAnswer(returned) {
		t.Errorf("twoSumFive failed, got:%v, want:%v", returned, twoSumFive.Answer)
	}

	if returned := twoSum(twoSumSix.Array, twoSumSix.Target); !twoSumSix.MatchesAnswer(returned) {
		t.Errorf("twoSumSix failed, got:%v, want:%v", returned, twoSumSix.Answer)
	}
}
