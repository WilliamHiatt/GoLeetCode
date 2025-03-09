package Stack

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {
	shortValidParentheses := ValidParenthesesBuilder("((()))", true)
	shortValidParentheses2 := ValidParenthesesBuilder("({[]})", true)
	longValidParentheses := ValidParenthesesBuilder("(()[{}](()))", true)
	longValidParentheses2 := ValidParenthesesBuilder("(((([[[[[{{{(({}))}}}]]]]]))))", true)

	shortInvalidParentheses := ValidParenthesesBuilder(")()", false)
	shortInvalidParentheses2 := ValidParenthesesBuilder("({(})", false)
	longInvalidParentheses := ValidParenthesesBuilder("({(([()])})", false)
	longInvalidParentheses2 := ValidParenthesesBuilder("{(([((((({", false)

	if result := IsValidParentheses(shortValidParentheses.InString); result != shortValidParentheses.Answer {
		t.Errorf("shortValidParentheses failed, got:%v, expected:%v", result, shortValidParentheses.Answer)
	}

	if result := IsValidParentheses(shortValidParentheses2.InString); result != shortValidParentheses2.Answer {
		t.Errorf("shortValidParentheses2 failed, got:%v, expected:%v", result, shortValidParentheses2.Answer)
	}

	if result := IsValidParentheses(longValidParentheses.InString); result != longValidParentheses.Answer {
		t.Errorf("longValidParentheses failed, got:%v, expected:%v", result, longValidParentheses.Answer)
	}

	if result := IsValidParentheses(longValidParentheses2.InString); result != longValidParentheses2.Answer {
		t.Errorf("longValidParentheses2 failed, got:%v, expected:%v", result, longValidParentheses2.Answer)
	}

	if result := IsValidParentheses(shortInvalidParentheses.InString); result != shortInvalidParentheses.Answer {
		t.Errorf("shortInvalidParentheses failed, got:%v, expected:%v", result, shortInvalidParentheses.Answer)
	}

	if result := IsValidParentheses(shortInvalidParentheses2.InString); result != shortInvalidParentheses2.Answer {
		t.Errorf("shortInvalidParentheses2 failed, got:%v, expected:%v", result, shortInvalidParentheses2.Answer)
	}

	if result := IsValidParentheses(longInvalidParentheses.InString); result != longInvalidParentheses.Answer {
		t.Errorf("longInvalidParentheses failed, got:%v, expected:%v", result, longInvalidParentheses.Answer)
	}

	if result := IsValidParentheses(longInvalidParentheses2.InString); result != longInvalidParentheses2.Answer {
		t.Errorf("longInvalidParentheses2 failed, got:%v, expected:%v", result, longInvalidParentheses2.Answer)
	}
}
