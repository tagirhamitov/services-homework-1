package types

import "fmt"

type Struct struct {
	Str   string
	Arr   []int32
	Dict  StringMap
	Int   int32
	Float float64
}

func NewStruct() Struct {
	s := Struct{
		Str:   "hello, wordld!",
		Arr:   nil,
		Dict:  make(map[string]string),
		Int:   42,
		Float: 3.141592,
	}
	for i := int32(0); i < 100; i++ {
		s.Arr = append(s.Arr, i)

		key := fmt.Sprintf("key%v", i)
		value := fmt.Sprintf("value%v", i)
		s.Dict[key] = value
	}
	return s
}
