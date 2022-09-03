package goassert

import "testing"

func Test_EqualShouldPass_WhenActualMatchesExpected(t *testing.T) {
	tester := new(testing.T)

	actual := "expected value"
	expected := "expected value"

	Equal(tester, expected, actual)

	if tester.Failed() {
		t.Error("Equal did not pass when actual matched expected")
	}
}

func Test_EqualShouldFail_WhenActualDoesNotMatchExpected(t *testing.T) {
	tester := new(testing.T)

	actual := "actual value"
	expected := "expected value"

	Equal(tester, expected, actual)

	if !tester.Failed() {
		t.Error("Equal did not fail when actual did not match expected")
	}
}

func Test_NotEqualShouldPass_WhenActualDoesNotMatchExpected(t *testing.T) {
	tester := new(testing.T)

	actual := "actual value"
	expected := "expected value"

	NotEqual(tester, expected, actual)

	if tester.Failed() {
		t.Error("NotEqual did not pass when actual did not match expected")
	}
}

func Test_NotEqualShouldFail_WhenActualMatchesExpected(t *testing.T) {
	tester := new(testing.T)

	actual := "expected value"
	expected := "expected value"

	NotEqual(tester, expected, actual)

	if !tester.Failed() {
		t.Error("NotEqual did not fail when actual matched expected")
	}
}

func Test_DeepEqualShouldPass_WhenActualMatchesExpected_GivenPrimitiveTypes(t *testing.T) {
	tester := new(testing.T)

	actual := 10
	expected := 10

	DeepEqual(tester, expected, actual)

	if tester.Failed() {
		t.Error("DeepEqual did not pass when actual matched expected given primitive types")
	}
}

func Test_DeepEqualShouldPass_WhenActualMatchesExpected_GivenStructs(t *testing.T) {
	tester := new(testing.T)

	actual := *newMockStruct(10)
	expected := *newMockStruct(10)

	DeepEqual(tester, expected, actual)

	if tester.Failed() {
		t.Error("DeepEqual did not pass when actual matched expected given structs")
	}
}

func Test_DeepEqualShouldPass_WhenActualMatchesExpected_GivenPointersToStruct(t *testing.T) {
	tester := new(testing.T)

	actual := newMockStruct(10)
	expected := newMockStruct(10)

	DeepEqual(tester, expected, actual)

	if tester.Failed() {
		t.Error("DeepEqual did not pass when actual matched expected given pointers to struct")
	}
}

func Test_DeepEqualShouldPass_WhenActualMatchesExpected_GivenSlicesEqualInValuesAndOrder(t *testing.T) {
	tester := new(testing.T)

	expectedSlice := []int{3, 10, 5, 16}
	actualSlice := []int{3, 10, 5, 16}

	DeepEqual(tester, expectedSlice, actualSlice)

	if tester.Failed() {
		t.Error("DeepEqual did not pass when given two slices with same elements in the same order")
	}
}

func Test_DeepEqualShouldPass_WhenActualMatchesExpected_GivenSliceOfPointers(t *testing.T) {
	tester := new(testing.T)

	expectedSlice := []*mockStruct{newMockStruct(10), newMockStruct(16)}
	actualSlice := []*mockStruct{newMockStruct(10), newMockStruct(16)}

	DeepEqual(t, expectedSlice, actualSlice)

	if tester.Failed() {
		t.Error("DeepEqual did not pass when given two slices of pointers with the same elements in same order")
	}
}

func Test_DeepEqualShouldPass_WhenActualMatchesExpected_GivenMapsWithEqualKeysAndValues(t *testing.T) {
	tester := new(testing.T)

	expectedMap := map[int]*mockStruct{
		10: newMockStruct(10),
		5:  newMockStruct(5),
		16: newMockStruct(16),
	}
	actualMap := map[int]*mockStruct{
		10: newMockStruct(10),
		5:  newMockStruct(5),
		16: newMockStruct(16),
	}

	DeepEqual(tester, expectedMap, actualMap)

	if tester.Failed() {
		t.Error("DeepEqual did not pass when given two maps with same key-value pairs")
	}
}

func Test_DeepEqualShouldFail_WhenActualDoesNotMatchExpected_GivenPrimitiveTypes(t *testing.T) {
	tester := new(testing.T)

	actual := 10
	expected := 16

	DeepEqual(tester, expected, actual)

	if !tester.Failed() {
		t.Error("DeepEqual did not fail when actual did not match expected given primitive types")
	}
}

func Test_DeepEqualShouldFail_WhenActualDoesNotMatchExpected_GivenStructs(t *testing.T) {
	tester := new(testing.T)

	actual := *newMockStruct(10)
	expected := *newMockStruct(16)

	DeepEqual(tester, expected, actual)

	if !tester.Failed() {
		t.Error("DeepEqual did not fail when actual did not match expected given structs")
	}
}

func Test_DeepEqualShouldFail_WhenActualDoesNotMatchExpected_GivenPointersToStruct(t *testing.T) {
	tester := new(testing.T)

	actual := newMockStruct(10)
	expected := newMockStruct(16)

	DeepEqual(tester, expected, actual)

	if !tester.Failed() {
		t.Error("DeepEqual did not fail when actual did not match expected given pointers to struct")
	}
}

func Test_DeepEqualShouldFail_WhenActualDoesNotMatchExpected_GivenSlices(t *testing.T) {
	tester := new(testing.T)

	expectedSlice := []int{3, 16, 5, 7}
	actualSlice := []int{3, 10, 5, 16}

	DeepEqual(tester, expectedSlice, actualSlice)

	if !tester.Failed() {
		t.Error("DeepEqual did not fail when given two different slices")
	}
}

func Test_DeepEqualShouldFail_WhenActualDoesNotMatchExpected_GivenMaps(t *testing.T) {
	tester := new(testing.T)

	expectedMap := map[int]*mockStruct{
		10: newMockStruct(10),
		5:  newMockStruct(5),
		16: newMockStruct(16),
	}
	actualMap := map[int]*mockStruct{
		10: newMockStruct(16),
		5:  newMockStruct(5),
		16: newMockStruct(16),
	}

	DeepEqual(tester, expectedMap, actualMap)

	if !tester.Failed() {
		t.Error("DeepEqual did not fail when given two different maps")
	}
}

func Test_NotDeepEqualShouldPass_WhenActualDoesNotMatchExpected_GivenPrimitiveTypes(t *testing.T) {
	tester := new(testing.T)

	actual := 10
	expected := 16

	NotDeepEqual(tester, expected, actual)

	if tester.Failed() {
		t.Error("NotDeepEqual did not pass when actual did not match expected given primitive types")
	}
}

func Test_NotDeepEqualShouldPass_WhenActualDoesNotMatchExpected_GivenStructs(t *testing.T) {
	tester := new(testing.T)

	actual := *newMockStruct(10)
	expected := *newMockStruct(16)

	NotDeepEqual(tester, expected, actual)

	if tester.Failed() {
		t.Error("NotDeepEqual did not pass when actual did not match expected given structs")
	}
}

func Test_NotDeepEqualShouldPass_WhenActualDoesNotMatchExpected_GivenPointersToStruct(t *testing.T) {
	tester := new(testing.T)

	actual := newMockStruct(10)
	expected := newMockStruct(16)

	NotDeepEqual(tester, expected, actual)

	if tester.Failed() {
		t.Error("NotDeepEqual did not pass when actual did not match expected given pointers to struct")
	}
}

func Test_NotDeepEqualShouldPass_WhenActualDidNotMatchExpected_GivenSlicesOfValues(t *testing.T) {
	tester := new(testing.T)

	expectedSlice := []int{3, 10, 5, 16}
	actualSlice := []int{10, 16, 5, 16}

	NotDeepEqual(tester, expectedSlice, actualSlice)

	if tester.Failed() {
		t.Error("NotDeepEqual did not pass when given two different slices of values")
	}
}

func Test_NotDeepEqualShouldPass_WhenActualDidNotMatchExpected_GivenSliceOfPointers(t *testing.T) {
	tester := new(testing.T)

	expectedSlice := []*mockStruct{newMockStruct(10), newMockStruct(16)}
	actualSlice := []*mockStruct{newMockStruct(10), newMockStruct(5)}

	NotDeepEqual(t, expectedSlice, actualSlice)

	if tester.Failed() {
		t.Error("NotDeepEqual did not pass when given two different slices of pointers")
	}
}

func Test_NotDeepEqualShouldPass_WhenActualDidNotMatchExpected_GivenMaps(t *testing.T) {
	tester := new(testing.T)

	expectedMap := map[int]*mockStruct{
		10: newMockStruct(10),
		5:  newMockStruct(5),
		16: newMockStruct(16),
	}
	actualMap := map[int]*mockStruct{
		10: newMockStruct(10),
		5:  newMockStruct(16),
		16: newMockStruct(16),
	}

	NotDeepEqual(tester, expectedMap, actualMap)

	if tester.Failed() {
		t.Error("NotDeepEqual did not pass when given two different maps")
	}
}

func Test_NotDeepEqualShouldFail_WhenActualMatchesExpected_GivenPrimitiveTypes(t *testing.T) {
	tester := new(testing.T)

	actual := 10
	expected := 10

	NotDeepEqual(tester, expected, actual)

	if !tester.Failed() {
		t.Error("NotDeepEqual did not fail when actual matched expected given primitive types")
	}
}

func Test_NotDeepEqualShouldFail_WhenActualMatchesExpected_GivenStructs(t *testing.T) {
	tester := new(testing.T)

	actual := *newMockStruct(10)
	expected := *newMockStruct(10)

	NotDeepEqual(tester, expected, actual)

	if !tester.Failed() {
		t.Error("NotDeepEqual did not fail when actual matched expected given structs")
	}
}

func Test_NotDeepEqualShouldFail_WhenActualMatchesExpected_GivenPointersToStruct(t *testing.T) {
	tester := new(testing.T)

	actual := newMockStruct(10)
	expected := newMockStruct(10)

	NotDeepEqual(tester, expected, actual)

	if !tester.Failed() {
		t.Error("NotDeepEqual did not fail when actual matched expected given pointers to struct")
	}
}

func Test_NotDeepEqualShouldFail_WhenActualMatchesExpected_GivenSlices(t *testing.T) {
	tester := new(testing.T)

	expectedSlice := []int{3, 10, 5, 16}
	actualSlice := []int{3, 10, 5, 16}

	NotDeepEqual(tester, expectedSlice, actualSlice)

	if !tester.Failed() {
		t.Error("NotDeepEqual did not fail when given two equal slices")
	}
}

func Test_NotDeepEqualShouldFail_WhenActualMatchesExpected_GivenMaps(t *testing.T) {
	tester := new(testing.T)

	expectedMap := map[int]*mockStruct{
		10: newMockStruct(10),
		5:  newMockStruct(5),
		16: newMockStruct(16),
	}
	actualMap := map[int]*mockStruct{
		10: newMockStruct(10),
		5:  newMockStruct(5),
		16: newMockStruct(16),
	}

	NotDeepEqual(tester, expectedMap, actualMap)

	if !tester.Failed() {
		t.Error("NotDeepEqual did not fail when given two equal maps")
	}
}

func Test_NilShouldPass_GivenNilPointer(t *testing.T) {
	tester := new(testing.T)

	var valuePtr *string = nil

	Nil(tester, valuePtr)

	if tester.Failed() {
		t.Error("Nil did not pass when nil pointer was given")
	}
}

func Test_NilShouldPass_GivenNilMap(t *testing.T) {
	tester := new(testing.T)

	var nilMap map[int]int = nil

	Nil(tester, nilMap)

	if tester.Failed() {
		t.Error("Nil did not pass when nil map was given")
	}
}

func Test_NilShouldPass_GivenNilSlice(t *testing.T) {
	tester := new(testing.T)

	var nilSlice []int = nil

	Nil(tester, nilSlice)

	if tester.Failed() {
		t.Error("Nil did not pass when nil slice was given")
	}
}

func Test_NilShouldPass_GivenNilChannel(t *testing.T) {
	tester := new(testing.T)

	var nilChan chan int = nil

	Nil(tester, nilChan)

	if tester.Failed() {
		t.Error("Nil did not pass when nil channel was given")
	}
}

func Test_NilShouldPass_GivenNilFunction(t *testing.T) {
	tester := new(testing.T)

	var nilFunc func() = nil

	Nil(tester, nilFunc)

	if tester.Failed() {
		t.Error("Nil did not pass when nil function was given")
	}
}

func Test_NilShouldFail_GivenNonNilPointer(t *testing.T) {
	tester := new(testing.T)

	value := "value"
	valuePtr := &value

	Nil(tester, valuePtr)

	if !tester.Failed() {
		t.Error("Nil did not fail when non-nil value was given")
	}
}

func Test_NilShouldFail_GivenNonNilMap(t *testing.T) {
	tester := new(testing.T)

	nonNilMap := make(map[int]int, 0)

	Nil(tester, nonNilMap)

	if !tester.Failed() {
		t.Error("Nil did not fail when non-nil map was given")
	}
}

func Test_NilShouldFail_GivenNonNilSlice(t *testing.T) {
	tester := new(testing.T)

	nonNilSlice := []int{10}

	Nil(tester, nonNilSlice)

	if !tester.Failed() {
		t.Error("Nil did not fail when non-nil slice was given")
	}
}

func Test_NilShouldFail_GivenNonNilChannel(t *testing.T) {
	tester := new(testing.T)

	nonNilChan := make(chan int)

	Nil(tester, nonNilChan)

	if !tester.Failed() {
		t.Error("Nil did not fail when non-nil channel was given")
	}
}

func Test_NilShouldFail_GivenNonNilFunction(t *testing.T) {
	tester := new(testing.T)

	nonNilFunc := func() {}

	Nil(tester, nonNilFunc)

	if !tester.Failed() {
		t.Error("Nil did not fail when non-nil function was given")
	}
}

func Test_NotNilShouldPass_GivenNonNilPointer(t *testing.T) {
	tester := new(testing.T)

	value := "value"
	valuePtr := &value

	NotNil(tester, valuePtr)

	if tester.Failed() {
		t.Error("NotNil did not pass when non-nil pointer was given")
	}
}

func Test_NotNilShouldPass_GivenNonNilMap(t *testing.T) {
	tester := new(testing.T)

	nonNilMap := make(map[int]int, 0)

	NotNil(tester, nonNilMap)

	if tester.Failed() {
		t.Error("NotNil did not pass when non-nil map was given")
	}
}

func Test_NotNilShouldPass_GivenNonNilSlice(t *testing.T) {
	tester := new(testing.T)

	nonNilSlice := []int{10}

	NotNil(tester, nonNilSlice)

	if tester.Failed() {
		t.Error("NotNil did not pass when non-nil slice was given")
	}
}

func Test_NotNilShouldPass_GivenNonNilChan(t *testing.T) {
	tester := new(testing.T)

	nonNilChan := make(chan int)

	NotNil(tester, nonNilChan)

	if tester.Failed() {
		t.Error("NotNil did not pass when non-nil channel was given")
	}
}

func Test_NotNilShouldPass_GivenNonNilFunction(t *testing.T) {
	tester := new(testing.T)

	nonNilFunc := func() {}

	NotNil(tester, nonNilFunc)

	if tester.Failed() {
		t.Error("NotNil did not pass when non-nil function was given")
	}
}

func Test_NotNilShouldFail_GivenNilPointer(t *testing.T) {
	tester := new(testing.T)

	var valuePtr *string = nil

	NotNil(tester, valuePtr)

	if !tester.Failed() {
		t.Error("NotNil did not fail when nil pointer was given")
	}
}

func Test_NotNilShouldFail_GivenNilMap(t *testing.T) {
	tester := new(testing.T)

	var nilMap map[int]int = nil

	NotNil(tester, nilMap)

	if !tester.Failed() {
		t.Error("NotNil did not fail when nil map was given")
	}
}

func Test_NotNilShouldFail_GivenNilSlice(t *testing.T) {
	tester := new(testing.T)

	var nilSlice []int = nil

	NotNil(tester, nilSlice)

	if !tester.Failed() {
		t.Error("NotNil did not fail when nil slice was given")
	}
}

func Test_NotNilShouldFail_GivenNilChannel(t *testing.T) {
	tester := new(testing.T)

	var nilChan chan int = nil

	NotNil(tester, nilChan)

	if !tester.Failed() {
		t.Error("NotNil did not fail when nil channel was given")
	}
}

func Test_NotNilShouldFail_GivenNilFunction(t *testing.T) {
	tester := new(testing.T)

	var nilFunc func() = nil

	NotNil(tester, nilFunc)

	if !tester.Failed() {
		t.Error("NotNil did not fail when nil function was given")
	}
}

func Test_SimilarSliceShouldPass_WhenActualSliceMatchesExpected_InValues_ButNotInOrder(t *testing.T) {
	tester := new(testing.T)

	expectedSlice := []int{3, 16, 5, 5, 10}
	actualSlice := []int{3, 10, 5, 5, 16}

	SimilarSlice(tester, expectedSlice, actualSlice)

	if tester.Failed() {
		t.Error("SimilarSlice did not pass when given two slices with same values but in different order")
	}
}

func Test_SimilarSliceShouldPass_WhenActualSliceMatchesExpected_InValuesAndOrder(t *testing.T) {
	tester := new(testing.T)

	expectedSlice := []int{3, 10, 5, 16, 16}
	actualSlice := []int{3, 10, 5, 16, 16}

	SimilarSlice(tester, expectedSlice, actualSlice)

	if tester.Failed() {
		t.Error("SimilarSlice did not pass when given two slices with same values in the same order")
	}
}

func Test_SimilarSliceShouldFail_WhenActualSliceDoesNotMatchExpected_InLengths(t *testing.T) {
	tester := new(testing.T)

	expectedSlice := []int{3, 10, 5}
	actualSlice := []int{3, 10, 5, 8}

	SimilarSlice(tester, expectedSlice, actualSlice)

	if !tester.Failed() {
		t.Error("SimilarSlice did not fail when given two slices with different lengths")
	}
}

func Test_SimilarSliceShouldFail_WhenActualSliceDoesNotMatchExpected_InValues(t *testing.T) {
	tester := new(testing.T)

	expectedSlice := []int{3, 7, 5, 5, 16}
	actualSlice := []int{3, 10, 5, 5, 8}

	SimilarSlice(tester, expectedSlice, actualSlice)

	if !tester.Failed() {
		t.Error("SimilarSlice did not fail when given two slices with different values")
	}
}

func Test_NotSimilarSliceShouldPass_WhenActualAndExpectedSlicesHaveDifferentLengths(t *testing.T) {
	tester := new(testing.T)

	expectedSlice := []int{3, 7, 5, 16}
	actualSlice := []int{3, 10, 5, 5, 8}

	NotSimilarSlice(tester, expectedSlice, actualSlice)

	if tester.Failed() {
		t.Error("NotSimilarSlice did not pass given two slices with different lengths")
	}
}

func Test_NotSimilarSliceShouldPass_WhenActualAndExpectedSlicesHaveDifferentValues(t *testing.T) {
	tester := new(testing.T)

	expectedSlice := []int{3, 7, 5, 8, 16}
	actualSlice := []int{3, 10, 5, 5, 8}

	NotSimilarSlice(tester, expectedSlice, actualSlice)

	if tester.Failed() {
		t.Error("NotSimilarSlice did not pass given two slices with different values")
	}
}

func Test_NotSimilarSliceShouldFail_WhenActualAndExpectedSlicesHaveSameValues(t *testing.T) {
	tester := new(testing.T)

	expectedSlice := []int{3, 5, 5, 10, 16}
	actualSlice := []int{3, 10, 5, 5, 16}

	NotSimilarSlice(tester, expectedSlice, actualSlice)

	if !tester.Failed() {
		t.Error("NotSimilarSlice did not fail given two slices with same values")
	}
}
