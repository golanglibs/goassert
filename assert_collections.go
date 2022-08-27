package goassert

import "testing"

func EmptySlice[T any](t testing.TB, s []T) {
	t.Helper()

	if s == nil {
		t.Error("Expected empty slice but got nil")
		return
	}

	length := len(s)
	if length != 0 {
		t.Errorf("Expected empty slice but got slice with length %d", length)
	}
}

func NotEmptySlice[T any](t testing.TB, s []T) {
	t.Helper()

	if s == nil {
		t.Error("Expected empty slice but got nil")
		return
	}

	if len(s) == 0 {
		t.Error("Expected non- empty slice but got empty slice")
	}
}

func SliceLength[T any](t testing.TB, s []T, expectedLength int) {
	t.Helper()

	length := len(s)
	if length != expectedLength {
		t.Errorf("Expected slice to have length of %d but got %d", expectedLength, length)
	}
}

func SliceContains[K comparable](t testing.TB, s []K, element K) {
	t.Helper()

	if !sliceContains(s, element) {
		t.Errorf("Element %v could not be found in the slice %v", element, s)
	}
}

func SliceNotContains[K comparable](t testing.TB, s []K, element K) {
	t.Helper()

	if sliceContains(s, element) {
		t.Errorf("Element %v was not expected to be found in the slice %v", element, s)
	}
}

func sliceContains[K comparable](s []K, element K) bool {
	for _, v := range s {
		if v == element {
			return true
		}
	}

	return false
}

func EmptyMap[K comparable, V any](t testing.TB, m map[K]V) {
	t.Helper()

	if m == nil {
		t.Error("Expected empty map but got nil")
		return
	}

	length := len(m)
	if length != 0 {
		t.Errorf("Expected empty map but got map with length of %d", length)
	}
}

func NotEmptyMap[K comparable, V any](t testing.TB, m map[K]V) {
	t.Helper()

	if m == nil {
		t.Error("Expected non-empty map but got nil")
		return
	}

	if len(m) == 0 {
		t.Error("Expected non-empty map but got empty map")
	}
}

func MapLength[K comparable, V any](t testing.TB, m map[K]V, expectedLength int) {
	t.Helper()

	length := len(m)
	if length != expectedLength {
		t.Errorf("Expected map to have length of %d but got %d", expectedLength, length)
	}
}

func MapContains[K, V comparable](t testing.TB, m map[K]V, k K, v V) {
	t.Helper()

	actualValue, found := m[k]

	if !found {
		t.Errorf("Key %v was not found in the map", k)
		return
	}

	if v != actualValue {
		t.Errorf("Expected %v for key %v in the map but got %v", v, k, actualValue)
	}
}

func MapNotContains[K, V comparable](t testing.TB, m map[K]V, k K, v V) {
	t.Helper()

	value, found := m[k]

	if found && v == value {
		t.Errorf("Key %v and value %v was not expected to be found in the map", k, v)
	}
}
