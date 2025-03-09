package TwoPointers

import (
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	// https://leetcode.com/problems/valid-palindrome/description/
	singleWordValid := ValidPalindromeBuilder("racecar", true)
	singleWordValid2 := ValidPalindromeBuilder("radar", true)
	multiWordValid := ValidPalindromeBuilder("A man, a plan, a canal, Panama!", true)
	multiWordValid2 := ValidPalindromeBuilder("Was it a car or a cat I saw?", true)
	emptyValid := ValidPalindromeBuilder(" ", true)

	singleWordInvalid := ValidPalindromeBuilder("this", false)
	singleWordInvalid2 := ValidPalindromeBuilder("seahawks", false)
	multiWordInvalid := ValidPalindromeBuilder("The Seahawks are trading a lot of players", false)
	multiWordInvalid2 := ValidPalindromeBuilder("Call of Duty Warzone is Fun!", false)

	if returned := validPalindrome(singleWordValid2.InString); returned != singleWordValid2.Answer {
		t.Errorf("singleWordValid2 did not pass, got:%v, expected:%v", returned, singleWordValid2.Answer)
	}

	if returned := validPalindrome(singleWordValid.InString); returned != singleWordValid.Answer {
		t.Errorf("singleWorldValid did not pass, got:%v, expected:%v", returned, singleWordValid.Answer)
	}

	if returned := validPalindrome(multiWordValid.InString); returned != multiWordValid.Answer {
		t.Errorf("multiWordValid did not pass, got:%v, expected:%v", returned, multiWordValid.Answer)
	}

	if returned := validPalindrome(multiWordValid2.InString); returned != multiWordValid2.Answer {
		t.Errorf("multiWordValid2 did not pass, got:%v, expected:%v", returned, multiWordValid2.Answer)
	}

	if returned := validPalindrome(emptyValid.InString); returned != emptyValid.Answer {
		t.Errorf("emptyValid did not pass, got:%v, expected:%v", returned, emptyValid.Answer)
	}

	if returned := validPalindrome(singleWordInvalid.InString); returned != singleWordInvalid.Answer {
		t.Errorf("singleWordInvalid did not pass, got:%v, expected:%v", returned, singleWordInvalid.Answer)
	}

	if returned := validPalindrome(singleWordInvalid2.InString); returned != singleWordInvalid2.Answer {
		t.Errorf("singleWordInvalid2 did not pass, got:%v, expected:%v", returned, singleWordInvalid2.Answer)
	}

	if returned := validPalindrome(multiWordInvalid.InString); returned != multiWordInvalid.Answer {
		t.Errorf("multiWordInvalid did not pass, got:%v, expected:%v", returned, multiWordInvalid.Answer)
	}

	if returned := validPalindrome(multiWordInvalid2.InString); returned != multiWordInvalid2.Answer {
		t.Errorf("multiWordInvalid2 did not pass, got:%v, expected:%v", returned, multiWordInvalid2.Answer)
	}
}
