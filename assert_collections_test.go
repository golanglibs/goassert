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

func Test_SliceContainsShouldPass_GivenElementInGivenSlice(t *testing.T) {
	tester := new(testing.T)

	slice := []int{10, 5, 16, 3}
	element := 5

	SliceContains(tester, slice, element)

	if tester.Failed() {
		t.Error("SliceContains did not pass when given element is found within given slice")
	}
}

func Test_SliceContainsShouldFail_GivenElementNotInGivenSlice(t *testing.T) {
	tester := new(testing.T)

	slice := []int{10, 5, 16, 3}
	element := 7

	SliceContains(tester, slice, element)

	if !tester.Failed() {
		t.Error("SliceContains did not fail when given element is not found within given slice")
	}
}

func Test_SliceNotContainsShouldPass_GivenElementNotInGivenSlice(t *testing.T) {
	tester := new(testing.T)

	slice := []int{10, 5, 16, 3}
	element := 7

	SliceNotContains(tester, slice, element)

	if tester.Failed() {
		t.Error("SliceNotContains did not pass when given element is not found within given slice")
	}
}

func Test_SliceNotContainsShouldFail_GivenElementInGivenSlice(t *testing.T) {
	tester := new(testing.T)

	slice := []int{10, 5, 16, 3}
	element := 10

	SliceNotContains(tester, slice, element)

	if !tester.Failed() {
		t.Error("SliceNotContains did not fail when given element is not found within given slice")
	}
}

func Test_EmptyMapShouldPass_GivenEmptyMap(t *testing.T) {
	tester := new(testing.T)

	m := make(map[int]int, 0)

	EmptyMap(tester, m)

	if tester.Failed() {
		t.Error("EmptyMap did not pass given empty map")
	}
}

func Test_EmptyMapShouldFail_GivenNilEmptyMap(t *testing.T) {
	tester := new(testing.T)

	EmptyMap[int, int](tester, nil)

	if !tester.Failed() {
		t.Error("EmptyMap did not fail given nil map")
	}
}

func Test_EmptyMapShouldFail_GivenNonEmptyMap(t *testing.T) {
	tester := new(testing.T)

	m := map[int]int{
		10: 10,
		5:  5,
	}

	EmptyMap(tester, m)

	if !tester.Failed() {
		t.Error("EmptyMap did not fail given non-empty map")
	}
}

func Test_NotEmptyMapShouldPass_GivenNonEmptyMap(t *testing.T) {
	tester := new(testing.T)

	m := map[int]int{
		10: 5,
		16: 16,
	}

	NotEmptyMap(tester, m)

	if tester.Failed() {
		t.Error("NotEmptyMap did not pass given non-empty map")
	}
}

func Test_NotEmptyMapShouldFail_GivenNilMap(t *testing.T) {
	tester := new(testing.T)

	NotEmptyMap[int, string](tester, nil)

	if !tester.Failed() {
		t.Error("NotEmptyMap did not fail given nil map")
	}
}

func Test_NotEmptyMapShouldFail_GivenEmptyMap(t *testing.T) {
	tester := new(testing.T)

	m := make(map[int]int, 0)

	NotEmptyMap(tester, m)

	if !tester.Failed() {
		t.Error("NotEmptyMap did not fail given empty map")
	}
}

func Test_MapLengthShouldPass_IfGivenMapLengthMatchesGivenLength(t *testing.T) {
	tester := new(testing.T)

	m := map[int]int{
		10: 10,
		5:  5,
	}

	MapLength(tester, m, 2)

	if tester.Failed() {
		t.Error("MapLength did not pass when given map's length matches expected length")
	}
}

func Test_MapLengthShouldFail_IfGivenMapLengthDoesNotMatchGivenLength(t *testing.T) {
	tester := new(testing.T)

	m := map[int]int{
		10: 10,
		5:  5,
	}

	MapLength(tester, m, 10)

	if !tester.Failed() {
		t.Error("MapLength did not fail when given map's length does not match expected length")
	}
}

func Test_MapContainsShouldPass_GivenKeyValuePairInGivenMap(t *testing.T) {
	tester := new(testing.T)

	m := map[int]int{
		5:  5,
		10: 10,
		16: 16,
	}
	k, v := 10, 10

	MapContains(tester, m, k, v)

	if tester.Failed() {
		t.Error("MapContains did not pass when given key value pair is found in given map")
	}
}

func Test_MapContainsShouldFail_GivenKeyNotInGivenMap(t *testing.T) {
	tester := new(testing.T)

	m := map[int]int{
		5:  5,
		10: 10,
		16: 16,
	}
	k, v := 7, 10

	MapContains(tester, m, k, v)

	if !tester.Failed() {
		t.Error("MapContains did not fail when given key value pair is not found in given map")
	}
}

func Test_MapNotContainsShouldPass_GivenKeyValuePairNotInGivenMap(t *testing.T) {
	tester := new(testing.T)

	m := map[int]int{
		5:  5,
		10: 10,
		16: 16,
	}
	k, v := 3, 7

	MapNotContains(tester, m, k, v)

	if tester.Failed() {
		t.Error("MapContains did not pass when given key value pair is not found in given map")
	}
}

func Test_MapNotContainsShouldPass_GivenValueNotInGivenMap(t *testing.T) {
	tester := new(testing.T)

	m := map[int]int{
		5:  5,
		10: 10,
		16: 16,
	}
	k, v := 10, 7

	MapNotContains(tester, m, k, v)

	if tester.Failed() {
		t.Error("MapContains did not pass when given key value pair is not found in given map")
	}
}

func Test_MapNotContainsShouldFail_GivenKeyValuePairInGivenMap(t *testing.T) {
	tester := new(testing.T)

	m := map[int]int{
		5:  5,
		10: 10,
		16: 16,
	}
	k, v := 5, 5

	MapNotContains(tester, m, k, v)

	if !tester.Failed() {
		t.Error("MapContains did not fail when given key value pair is found in given map")
	}
}
