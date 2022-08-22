package goassert

import "testing"

func Test_PanicShouldPass_WhenGivenFuncPanics(t *testing.T) {
	tester := new(testing.T)

	Panic(tester, func() {
		panic("Error")
	})

	if tester.Failed() {
		t.Error("Panic did not pass when the given func panicked")
	}
}

func Test_PanicShouldFail_WhenGivenFuncDoesNotPanic(t *testing.T) {
	tester := new(testing.T)

	Panic(tester, func() {})

	if !tester.Failed() {
		t.Error("Panic did not fail when the given func did not panic")
	}
}

func Test_NotPanicShouldPass_WhenGivenFuncDoesNotPanic(t *testing.T) {
	tester := new(testing.T)

	NotPanic(tester, func() {})

	if tester.Failed() {
		t.Error("NotPanic did not pass when the given func did not panic")
	}
}

func Test_NotPanicShouldFail_WhenGivenFuncPanics(t *testing.T) {
	tester := new(testing.T)

	NotPanic(tester, func() {
		panic("Error")
	})

	if !tester.Failed() {
		t.Error("NotPanic did not fail when the given func panicked")
	}
}

func Test_PanicWithErrorShouldPass_WhenGivenFuncPanicsWithGivenError(t *testing.T) {
	tester := new(testing.T)

	error := "Error"
	PanicWithError(tester, error, func() {
		panic(error)
	})

	if tester.Failed() {
		t.Error("PanicWithError did not pass when the given func panicked with given error")
	}
}

func Test_PanicWithErrorShouldFail_WhenGivenFuncPanicsWithDifferentError(t *testing.T) {
	tester := new(testing.T)

	error := "Error"
	PanicWithError(tester, error, func() {
		panic("Different Error")
	})

	if !tester.Failed() {
		t.Error("PanicWithError did not fail when the given func panicked with different error")
	}
}

func Test_PanicWithErrorShouldFail_WhenGivenFuncDoesNotPanic(t *testing.T) {
	tester := new(testing.T)

	error := "Error"
	PanicWithError(tester, error, func() {})

	if !tester.Failed() {
		t.Error("PanicWithError did not fail when the given func did not panic")
	}
}

func Test_NotPanicWithErrorShouldPass_WhenGivenFuncDoesNotPanic(t *testing.T) {
	tester := new(testing.T)

	error := "Error"
	NotPanicWithError(tester, error, func() {})

	if tester.Failed() {
		t.Error("NotPanicWithError did not pass when the given func did not panic")
	}
}

func Test_NotPanicWithErrorShouldPass_WhenGivenFuncPanicsWithDifferentError(t *testing.T) {
	tester := new(testing.T)

	error := "Error"
	NotPanicWithError(tester, error, func() {
		panic("Different Error")
	})

	if tester.Failed() {
		t.Error("NotPanicWithError did not pass when the given func panicked with different error")
	}
}

func Test_NotPanicWithErrorShouldFail_WhenGivenFuncPanicsWithGivenError(t *testing.T) {
	tester := new(testing.T)

	error := "Error"
	NotPanicWithError(tester, error, func() {
		panic(error)
	})

	if !tester.Failed() {
		t.Error("NotPanicWithError did not fail when the given func panicked with given error")
	}
}
