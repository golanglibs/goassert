package goassert

import "testing"

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

func Test_MapContainsKeyShouldPass_IfGivenKeyExistsInGivenMap(t *testing.T) {
	tester := new(testing.T)

	m := map[int]int{
		5:  5,
		10: 10,
		16: 16,
	}
	k := 10

	MapContainsKey(tester, m, k)

	if tester.Failed() {
		t.Error("MapContainsKey did not pass when given key is found in given map")
	}
}

func Test_MapContainsKeyShouldFail_IfGivenKeyDoesNotExistInGivenMap(t *testing.T) {
	tester := new(testing.T)

	m := map[int]int{
		5:  5,
		10: 10,
		16: 16,
	}
	k := 7

	MapContainsKey(tester, m, k)

	if !tester.Failed() {
		t.Error("MapContainsKey did not fail when given key is not found in given map")
	}
}

func Test_MapNotContainsKeyShouldPass_IfGivenKeyDoesNotExistInGivenMap(t *testing.T) {
	tester := new(testing.T)

	m := map[int]int{
		5:  5,
		10: 10,
		16: 16,
	}
	k := 7

	MapNotContainsKey(tester, m, k)

	if tester.Failed() {
		t.Error("MapNotContainsKey did not pass when given key is not found in given map")
	}
}

func Test_MapNotContainsKeyShouldFail_IfGivenKeyExistsInGivenMap(t *testing.T) {
	tester := new(testing.T)

	m := map[int]int{
		5:  5,
		10: 10,
		16: 16,
	}
	k := 10

	MapNotContainsKey(tester, m, k)

	if !tester.Failed() {
		t.Error("MapNotContainsKey did not pass when given key is found in given map")
	}
}

func Test_MapContainsShouldPass_IfGivenKeyValuePairExistsInGivenMap(t *testing.T) {
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

func Test_MapContainsShouldFail_IfGivenKeyDoesNotExistInGivenMap(t *testing.T) {
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

func Test_MapNotContainsShouldPass_IfGivenKeyValuePairDoesNotExistInGivenMap(t *testing.T) {
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

func Test_MapNotContainsShouldPass_IfGivenValueDoesNotExistInGivenMap(t *testing.T) {
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

func Test_MapNotContainsShouldFail_IfGivenKeyValuePairExistsInGivenMap(t *testing.T) {
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
