package goassert

import "testing"

/*
Asserts that the given value is true
*/
func True(t testing.TB, assertion bool) {
	t.Helper()

	if !assertion {
		t.Error(inequalityMsg(true, false))
	}
}

/*
Asserts that the given value is false
*/
func False(t testing.TB, assertion bool) {
	t.Helper()

	if assertion {
		t.Error(inequalityMsg(false, true))
	}
}
