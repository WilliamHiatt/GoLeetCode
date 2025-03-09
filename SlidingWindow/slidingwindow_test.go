package SlidingWindow

import "testing"

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	bttbassTestCase := BTTBASSBuilder([]int{7, 1, 5, 3, 6, 4}, 5)
	bttbassTestCase2 := BTTBASSBuilder([]int{7, 6, 4, 3, 1}, 0)
	bttbassTestCase3 := BTTBASSBuilder([]int{2, 4, 1, 5, 3, 6, 8, 3}, 7)
	bttbassTestCase4 := BTTBASSBuilder([]int{3, 2, 6, 5, 0, 3}, 4)
	bttbassTestCase5 := BTTBASSBuilder([]int{5}, 0)

	if answer := GetBTTBASS(bttbassTestCase.IncomingArr); answer != bttbassTestCase.Answer {
		t.Errorf("bttbassTestCase: Got:%v, Expected:%v", answer, bttbassTestCase.Answer)
	}

	if answer := GetBTTBASS(bttbassTestCase2.IncomingArr); answer != bttbassTestCase2.Answer {
		t.Errorf("bttbassTestCase2: Got:%v, Expected:%v", answer, bttbassTestCase2.Answer)
	}

	if answer := GetBTTBASS(bttbassTestCase3.IncomingArr); answer != bttbassTestCase3.Answer {
		t.Errorf("bttbassTestCase3: Got:%v, Expected:%v", answer, bttbassTestCase3.Answer)
	}

	if answer := GetBTTBASS(bttbassTestCase4.IncomingArr); answer != bttbassTestCase4.Answer {
		t.Errorf("bttbassTestCase4: Got:%v, Expected:%v", answer, bttbassTestCase4.Answer)
	}

	if answer := GetBTTBASS(bttbassTestCase5.IncomingArr); answer != bttbassTestCase5.Answer {
		t.Errorf("bttbassTestCase5: Got:%v, Expected:%v", answer, bttbassTestCase5.Answer)
	}
}
