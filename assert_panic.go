package goassert

import (
	"reflect"
	"testing"
)

/*
Asserts that the given function panics
*/
func Panic(t testing.TB, underTest func()) {
	t.Helper()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic but there was no panic")
		}
	}()

	underTest()
}

/*
Asserts that the given function does not panic
*/
func NotPanic(t testing.TB, underTest func()) {
	t.Helper()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected no panic but there was panic: %v", r)
		}
	}()

	underTest()
}

/*
Asserts that the given function panics with the specified error.
The actual error must be of the same type and value as the given error
*/
func PanicWithError[T any](t testing.TB, expectedError T, underTest func()) {
	t.Helper()

	defer func() {
		r := recover()
		if r == nil {
			t.Error("Expected panic but there was no panic")
			return
		}

		if !reflect.DeepEqual(expectedError, r) {
			t.Errorf("Expected panic with %v error but got %v error", expectedError, r)
		}
	}()

	underTest()
}

/*
Asserts that the given function does not panic with the specified error.
The assertion succeeds if the given function does not panic or panics with a different error
*/
func NotPanicWithError[T any](t testing.TB, expectedError T, underTest func()) {
	t.Helper()

	defer func() {
		r := recover()
		if r == nil {
			return
		}

		if reflect.DeepEqual(expectedError, r) {
			t.Errorf("Expected panic with different error than %v error", expectedError)
		}
	}()

	underTest()
}
