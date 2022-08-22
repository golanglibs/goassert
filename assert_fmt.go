package goassert

import "fmt"

var inqualityMsgTemplate string = "Expected: %v. Actual: %v"
var equalityMsgTemplate string = "Expected to not equal: %v"

func inequalityMsg[T any](expected T, actual T) string {
	return fmt.Sprintf(inqualityMsgTemplate, expected, actual)
}

func equalityMsg[T any](expected T) string {
	return fmt.Sprintf(equalityMsgTemplate, expected)
}
