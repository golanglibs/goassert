package goassert

import (
	"reflect"
	"testing"
)

/*
Asserts that the two given values are equal. The given values must be [comparable]
*/
func Equal[K comparable](t testing.TB, expected K, actual K) {
	t.Helper()

	if actual != expected {
		t.Error(inequalityMsg(expected, actual))
	}
}

/*
Asserts that the two given values are not equal. The given values must be [comparable]
*/
func NotEqual[K comparable](t testing.TB, expected K, actual K) {
	t.Helper()

	if actual == expected {
		t.Error(equalityMsg(expected))
	}
}

/*
Asserts that the two given valus are deeply equal. Internally uses reflect.DeepEqual.
The equality of arrays, slices and maps can be asserted with this method
*/
func DeepEqual[T any](t testing.TB, expected T, actual T) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		t.Error(inequalityMsg(expected, actual))
	}
}

/*
Asserts that the two given valus are not deeply equal. Internally uses reflect.DeepEqual.
The inequality of arrays, slices and maps can be asserted with this method
*/
func NotDeepEqual[T any](t testing.TB, expected T, actual T) {
	t.Helper()

	if reflect.DeepEqual(expected, actual) {
		t.Error(equalityMsg(expected))
	}
}

/*
Asserts that the given value is nil
*/
func Nil(t testing.TB, actual interface{}) {
	t.Helper()

	if !isNil(actual) {
		t.Error(inequalityMsg(nil, actual))
	}
}

/*
Asserts that the given value is not nil
*/
func NotNil(t testing.TB, actual interface{}) {
	t.Helper()

	if isNil(actual) {
		t.Error(equalityMsg("nil"))
	}
}

func isNil(value interface{}) bool {
	if value == nil {
		return true
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func:
		return reflect.ValueOf(value).IsNil()
	}

	return false
}

/*
Asserts that the two given slices have the same values in any order. The elements must be [comparable]
*/
func SimilarSlice[T any](t testing.TB, expected []T, actual []T) {
	t.Helper()

	if !areSimilarSlices(expected, actual) {
		t.Error(inequalityMsg(expected, actual))
	}
}

/*
Asserts that the two given slices does not have the same values. The elements must be [comparable]
*/
func NotSimilarSlice[T any](t testing.TB, expected []T, actual []T) {
	t.Helper()

	if areSimilarSlices(expected, actual) {
		t.Error(equalityMsg(expected))
	}
}

func areSimilarSlices[T any](expected []T, actual []T) bool {
	expectedLength := len(expected)
	actualLength := len(actual)
	if actualLength != expectedLength {
		return false
	}

	matchedIndexMap := make(map[int]interface{}, expectedLength)
	for _, expectedValue := range expected {
		for j, actualValue := range actual {
			if !reflect.DeepEqual(expectedValue, actualValue) {
				continue
			}

			if _, alreadyMatched := matchedIndexMap[j]; !alreadyMatched {
				matchedIndexMap[j] = nil
				break
			}
		}
	}

	return len(matchedIndexMap) == expectedLength
}
