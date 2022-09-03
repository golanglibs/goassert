# goassert

`goassert` is a simple assertion library for go/golang.
It does not try to do anything fancy and keeps things as simple as possible.
Its primary goal is to make assertions as short and simple as possible.
It also takes advantage of go generics as much as possible to be able to handle as many types as possible
with as fewest assertion methods as possible

## Installation
```bash
go get -t -u github.com/golanglibs/goassert@latest
```

## Usage
```go
package yourpackage

import (
	"testing",

	"github.com/golanglibs/goassert"
)

func Test_1Plus1ShouldEqual2(t *testing.T) {
	actual := 1 + 1
	expected := 2

	goassert.Equal(t, expected, actual)
	// on assertion error
	// --- FAIL: Test_1Plus1ShouldEqual2 (0.00s)
	// module_test.go: 11: Expected 2. Actual: 1
}
```

## Available Assertions

### Truth
* `True` - asserts the value is true
* `False` - asserts the value is false

### Equality
* `Equal` - asserts two values are equal. Values must be comparable
* `NotEqual` - asserts two values are not equal. Values must be comparable
* `Nil` - asserts the value is nil
* `NotNil` - asserts the value is not nil
* `SimilarSlice` - asserts two slices have the same values in any order. The elements must be comparable
* `NotSimilarSlice` - asserts two slices do not have the same values. The elements must be comparable

### Collection
* `EmptySlice` - asserts the slice is empty. The assertion will fail if the slice is nil
* `NotEmptySlice` - asserts the slice is not nil or empty
* `SliceLength` - asserts the slice has the specified length
* `SliceContains` - asserts the slice contains the specified value. The value must be comparable
* `SliceNotContains` - asserts the slice does not contain the specified value. The value must be comparable
* `EmptyMap` - asserts the map is empty. The assertion will fail if the map is nil
* `NotEmptyMap` - asserts the map is not nil or empty
* `MapLength` - asserts the map has the specified length
* `MapContainsKey` - asserts the map contains the specified key. Key must be comparable
* `MapNotContainsKey` - asserts the map does not contain the specified key. Key must be comparable
* `MapContains` - asserts the map contains the specified key-value pair. Key and value must be comparable
* `MapNotContains` - asserts the map does not contain the specified key-value pair. Key and value must be comparable

### Panic
* `Panic` - asserts given function panics
* `NotPanic` - asserts given function does not panic
* `PanicWithError` - asserts given function panics with the specified error
* `NotPanicWithError` - asserts given functoin does not panic with the specified error
