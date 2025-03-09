package arraysandhashing

type DuplicatesQuestion struct {
	ArrayOne []int
	Answer   bool
}

func NewDuplicateQuestion(arrayOne []int, answer bool) *DuplicatesQuestion {
	return &DuplicatesQuestion{
		ArrayOne: arrayOne,
		Answer:   answer,
	}
}

type ValidAnagram struct {
	FirstString  string
	SecondString string
	Answer       bool
}

func NewValidAnagram(firstString string, secondString string, answer bool) *ValidAnagram {
	return &ValidAnagram{
		FirstString:  firstString,
		SecondString: secondString,
		Answer:       answer,
	}
}

type TwoSum struct {
	Array  []int
	Target int
	Answer []int
}

func (ts TwoSum) MatchesAnswer(in []int) bool {
	if len(ts.Answer) != len(in) {
		return false
	}

	for i := 0; i < len(in); i++ {
		if ts.Answer[i] != in[i] {
			return false
		}
	}

	return true
}

func TwoSumQuestion(array []int, target int, answer []int) *TwoSum {
	return &TwoSum{
		Array:  array,
		Target: target,
		Answer: answer,
	}
}
