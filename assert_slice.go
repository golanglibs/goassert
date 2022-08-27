package goassert

import "testing"

/*
Asserts that the given slice is empty. The assertion fails if the given slice is nil
*/
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

/*
Asserts that the given slice is not nil or empty
*/
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

/*
Asserts that the given slice has length equal to the specified expected length
*/
func SliceLength[T any](t testing.TB, s []T, expectedLength int) {
	t.Helper()

	length := len(s)
	if length != expectedLength {
		t.Errorf("Expected slice to have length of %d but got %d", expectedLength, length)
	}
}

/*
Asserts that the given slice contains the given element. The element must be [comparable]
*/
func SliceContains[K comparable](t testing.TB, s []K, element K) {
	t.Helper()

	if !sliceContains(s, element) {
		t.Errorf("Element %v could not be found in the slice %v", element, s)
	}
}

/*
Asserts that the given slice does not contain the given element. The element must be [comparable]
*/
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
