package goassert

type mockStruct struct {
	Prop int
}

func newMockStruct(prop int) *mockStruct {
	return &mockStruct{
		Prop: prop,
	}
}
