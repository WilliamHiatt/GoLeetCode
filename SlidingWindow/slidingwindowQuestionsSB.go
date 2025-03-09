package SlidingWindow

type BTTBASS struct {
	IncomingArr []int
	Answer      int
}

func BTTBASSBuilder(incomingArr []int, answer int) *BTTBASS {
	return &BTTBASS{
		IncomingArr: incomingArr,
		Answer:      answer,
	}
}
