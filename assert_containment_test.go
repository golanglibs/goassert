package goassert

import "testing"

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
