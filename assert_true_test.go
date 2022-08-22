package goassert

import "testing"

func Test_TrueShouldPass_GivenFalseAssertion(t *testing.T) {
	tester := new(testing.T)
	True(tester, true)

	if tester.Failed() {
		t.Error("True did not pass given true assertion")
	}
}

func Test_TrueShouldFail_GivenFalseAssertion(t *testing.T) {
	tester := new(testing.T)
	True(tester, false)

	if !tester.Failed() {
		t.Error("True did not fail given false assertion")
	}
}

func Test_FalseShouldPass_GivenFalseAssertion(t *testing.T) {
	tester := new(testing.T)
	False(tester, false)

	if tester.Failed() {
		t.Error("False did not pass given false assertion")
	}
}

func Test_FalseShouldFail_GivenFalseAssertion(t *testing.T) {
	tester := new(testing.T)
	False(tester, true)

	if !tester.Failed() {
		t.Error("False did not fail given true assertion")
	}
}
