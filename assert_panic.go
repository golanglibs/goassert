package goassert

import (
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
func PanicWithError[K comparable](t testing.TB, expectedError K, underTest func()) {
	t.Helper()

	defer func() {
		r := recover()
		if r == nil {
			t.Error("Expected panic but there was no panic")
			return
		}

		error, converted := r.(K)
		if converted {
			Equal(t, expectedError, error)
		} else {
			t.Errorf(
				"Panic threw an unexpected error type. Expected Type: %T. Actual Type: %T",
				expectedError,
				r,
			)
		}
	}()

	underTest()
}

/*
Asserts that the given function does not panic with the specified error.
The assertion succeeds if the given function does not panic or panics with a different error
*/
func NotPanicWithError[K comparable](t testing.TB, expectedError K, underTest func()) {
	t.Helper()

	defer func() {
		r := recover()
		if r == nil {
			return
		}

		error, converted := r.(K)
		if converted {
			NotEqual(t, expectedError, error)
		}
	}()

	underTest()
}
