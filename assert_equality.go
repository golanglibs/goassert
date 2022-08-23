package goassert

import (
	"reflect"
	"testing"
)

func Equal[K comparable](t testing.TB, expected K, actual K) {
	t.Helper()

	if actual != expected {
		t.Error(inequalityMsg(expected, actual))
	}
}

func NotEqual[K comparable](t testing.TB, expected K, actual K) {
	t.Helper()

	if actual == expected {
		t.Error(equalityMsg(expected))
	}
}

func Nil(t testing.TB, actual interface{}) {
	t.Helper()

	if !isNil(actual) {
		t.Error(inequalityMsg(nil, actual))
	}
}

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
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(value).IsNil()
	}

	return false
}

func EqualSlice[K comparable](t testing.TB, expected []K, actual []K) {
	t.Helper()

	if !areEqualSlices(expected, actual) {
		t.Error(inequalityMsg(expected, actual))
	}
}

func NotEqualSlice[K comparable](t testing.TB, expected []K, actual []K) {
	t.Helper()

	if areEqualSlices(expected, actual) {
		t.Error(equalityMsg(expected))
	}
}

func areEqualSlices[K comparable](expected []K, actual []K) bool {
	expectedLength := len(expected)
	actualLength := len(actual)
	if actualLength != expectedLength {
		return false
	}

	for i := 0; i < expectedLength; i++ {
		if actual[i] != expected[i] {
			return false
		}
	}

	return true
}

func SimilarSlice[K comparable](t testing.TB, expected []K, actual []K) {
	t.Helper()

	if !areSimilarSlices(expected, actual) {
		t.Error(inequalityMsg(expected, actual))
	}
}

func NotSimilarSlice[K comparable](t testing.TB, expected []K, actual []K) {
	t.Helper()

	if areSimilarSlices(expected, actual) {
		t.Error(equalityMsg(expected))
	}
}

func areSimilarSlices[K comparable](expected []K, actual []K) bool {
	expectedLength := len(expected)
	actualLength := len(actual)
	if actualLength != expectedLength {
		return false
	}

	expectedElementsMap := make(map[K]int, expectedLength)
	for _, v := range expected {
		expectedElementsMap[v] += 1
	}

	actualElementsMap := make(map[K]int, actualLength)
	for _, v := range actual {
		actualElementsMap[v] += 1
	}

	return isEqualMap(expectedElementsMap, actualElementsMap)
}

func EqualMap[K, V comparable](t testing.TB, expected map[K]V, actual map[K]V) {
	t.Helper()

	if !isEqualMap(expected, actual) {
		t.Error(inequalityMsg(expected, actual))
	}
}

func NotEqualMap[K, V comparable](t testing.TB, expected map[K]V, actual map[K]V) {
	t.Helper()

	if isEqualMap(expected, actual) {
		t.Error(equalityMsg(expected))
	}
}

func isEqualMap[K, V comparable](expected map[K]V, actual map[K]V) bool {
	expectedLength := len(expected)
	actualLength := len(actual)
	if actualLength != expectedLength {
		return false
	}

	for actualKey, actualValue := range actual {
		expectedValue, found := expected[actualKey]
		if !found || actualValue != expectedValue {
			return false
		}
	}

	return true
}
