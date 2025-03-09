package BinarySearch

type BSQuestion struct {
	Input  []int
	Target int
	Answer int
}

func BSQuestionBuilder(input []int, target int, answer int) *BSQuestion {
	return &BSQuestion{
		Input:  input,
		Target: target,
		Answer: answer,
	}
}
