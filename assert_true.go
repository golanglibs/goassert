package goassert

import "testing"

func True(t testing.TB, assertion bool) {
	if !assertion {
		t.Error(inequalityMsg(true, false))
	}
}

func False(t testing.TB, assertion bool) {
	if assertion {
		t.Error(inequalityMsg(false, true))
	}
}
