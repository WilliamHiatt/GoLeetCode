package BinarySearch

import "testing"

func TestBinarySearch(t *testing.T) {
	validBinarySearch := BSQuestionBuilder([]int{3, 5, 7, 10, 22, 23, 25, 60, 61}, 10, 3)
	validBinarySearch2 := BSQuestionBuilder([]int{-11, -3, -1, 0, 1, 4, 6, 8, 11}, -11, 0)
	validBinarySearch3 := BSQuestionBuilder([]int{-11, -3, -1, 0, 1, 4, 6, 8, 11}, 11, 8)
	validBinarySearch4 := BSQuestionBuilder([]int{-11, -3, -1, 0, 1, 4, 6, 8, 11}, 0, 3)

	inValidBinarySearch := BSQuestionBuilder([]int{-11, -3, -1, 0, 1, 4, 6, 8, 11}, -2, -1)
	inValidBinarySearch2 := BSQuestionBuilder([]int{-11, -3, -1, 0, 1, 4, 6, 8, 11}, 2, -1)

	if answer := BSearch(validBinarySearch.Input, validBinarySearch.Target); answer != validBinarySearch.Answer {
		t.Errorf("validBinarySearch: got:%v, expected:%v", answer, validBinarySearch)
	}

	if answer := BSearch(validBinarySearch2.Input, validBinarySearch2.Target); answer != validBinarySearch2.Answer {
		t.Errorf("validBinarySearch2: got:%v, expected:%v", answer, validBinarySearch2)
	}

	if answer := BSearch(validBinarySearch3.Input, validBinarySearch3.Target); answer != validBinarySearch3.Answer {
		t.Errorf("validBinarySearch3: got:%v, expected:%v", answer, validBinarySearch3)
	}

	if answer := BSearch(validBinarySearch4.Input, validBinarySearch4.Target); answer != validBinarySearch4.Answer {
		t.Errorf("validBinarySearch4: got:%v, expected:%v", answer, validBinarySearch4)
	}

	if answer := BSearch(inValidBinarySearch.Input, inValidBinarySearch.Target); answer != inValidBinarySearch.Answer {
		t.Errorf("inValidBinarySearch: got:%v, expected:%v", answer, inValidBinarySearch)
	}

	if answer := BSearch(inValidBinarySearch2.Input, inValidBinarySearch2.Target); answer != inValidBinarySearch2.Answer {
		t.Errorf("inValidBinarySearch2: got:%v, expected:%v", answer, validBinarySearch)
	}
}
