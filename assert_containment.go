package goassert

import "testing"

func SliceContains[K comparable](t *testing.T, s []K, element K) {
	if !sliceContains(s, element) {
		t.Errorf("Element %v could not be found in the slice %v", element, s)
	}
}

func SliceNotContains[K comparable](t *testing.T, s []K, element K) {
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

func MapContains[K, V comparable](t *testing.T, m map[K]V, k K, v V) {
	actualValue, found := m[k]

	if !found {
		t.Errorf("Key %v was not found in the map", k)
		return
	}

	if v != actualValue {
		t.Errorf("Expected %v for key %v in the map but got %v", v, k, actualValue)
	}
}

func MapNotContains[K, V comparable](t *testing.T, m map[K]V, k K, v V) {
	value, found := m[k]

	if found && v == value {
		t.Errorf("Key %v and value %v was not expected to be found in the map", k, v)
	}
}
