package goassert

import "testing"

func Test_EmptySliceShouldPass_GivenEmptySlice(t *testing.T) {
	tester := new(testing.T)

	slice := make([]int, 0)

	EmptySlice(tester, slice)

	if tester.Failed() {
		t.Error("EmptySlice did not pass when empty slice was given")
	}
}

func Test_EmptySliceShouldFail_GivenNilSlice(t *testing.T) {
	tester := new(testing.T)

	EmptySlice[int](tester, nil)

	if !tester.Failed() {
		t.Error("EmptySlice did not fail when nil slice was given")
	}
}

func Test_EmptySliceShouldFail_GivenNonEmptySlice(t *testing.T) {
	tester := new(testing.T)

	slice := []int{10, 3, 5}

	EmptySlice(tester, slice)

	if !tester.Failed() {
		t.Error("EmptySlice did not fail when non-empty slice was given")
	}
}

func Test_NotEmptySliceShouldPass_GivenNonEmptySlice(t *testing.T) {
	tester := new(testing.T)

	slice := []int{10, 16}

	NotEmptySlice(tester, slice)

	if tester.Failed() {
		t.Error("NotEmptySlice did not pass when non-empty slice was given")
	}
}

func Test_NotEmptySliceShouldFail_GivenNilSlice(t *testing.T) {
	tester := new(testing.T)

	NotEmptySlice[string](tester, nil)

	if !tester.Failed() {
		t.Error("NotEmptySlice did not fail when nil slice was given")
	}
}

func Test_NotEmptySliceShouldFail_GivenEmptySlice(t *testing.T) {
	tester := new(testing.T)

	slice := make([]int, 0)

	NotEmptySlice(tester, slice)

	if !tester.Failed() {
		t.Error("NotEmptySlice did not fail when empty slice was given")
	}
}

func Test_SliceLengthShouldPass_IfGivenSliceLengthMatchesGivenLength(t *testing.T) {
	tester := new(testing.T)

	slice := []int{10, 3, 5}

	SliceLength(tester, slice, 3)

	if tester.Failed() {
		t.Error("SliceLength did not pass when given slice's length matches the expected length")
	}
}

func Test_SliceLengthShouldFail_IfGivenSliceLengthDoesNotMatchGivenLength(t *testing.T) {
	tester := new(testing.T)

	slice := []int{10, 3, 5}

	SliceLength(tester, slice, 5)

	if !tester.Failed() {
		t.Error("SliceLength did not fail when given slice's length does not match the expected length")
	}
}

func Test_SliceContainsShouldPass_IfGivenElementExistsInGivenSlice(t *testing.T) {
	tester := new(testing.T)

	slice := []int{10, 5, 16, 3}
	element := 5

	SliceContains(tester, slice, element)

	if tester.Failed() {
		t.Error("SliceContains did not pass when given element is found within given slice")
	}
}

func Test_SliceContainsShouldFail_IfGivenElementDoesNotExistInGivenSlice(t *testing.T) {
	tester := new(testing.T)

	slice := []int{10, 5, 16, 3}
	element := 7

	SliceContains(tester, slice, element)

	if !tester.Failed() {
		t.Error("SliceContains did not fail when given element is not found within given slice")
	}
}

func Test_SliceNotContainsShouldPass_IfGivenElementDoesNotExistInGivenSlice(t *testing.T) {
	tester := new(testing.T)

	slice := []int{10, 5, 16, 3}
	element := 7

	SliceNotContains(tester, slice, element)

	if tester.Failed() {
		t.Error("SliceNotContains did not pass when given element is not found within given slice")
	}
}

func Test_SliceNotContainsShouldFail_IfGivenElementExistsInGivenSlice(t *testing.T) {
	tester := new(testing.T)

	slice := []int{10, 5, 16, 3}
	element := 10

	SliceNotContains(tester, slice, element)

	if !tester.Failed() {
		t.Error("SliceNotContains did not fail when given element is not found within given slice")
	}
}
