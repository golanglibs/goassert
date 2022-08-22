package goassert

import (
	"testing"
)

func Panic(t *testing.T, underTest func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic but there was no panic")
		}
	}()

	underTest()
}

func NotPanic(t *testing.T, underTest func()) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected no panic but there was panic: %v", r)
		}
	}()

	underTest()
}

func PanicWithError[K comparable](t *testing.T, expectedError K, underTest func()) {
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

func NotPanicWithError[K comparable](t *testing.T, expectedError K, underTest func()) {
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
