package value

import (
	"fmt"
	"sort"
)

// FiberValues define the fiber's key / values
var FiberValues []*InstructionValue

// InstructionValue Define a key / value of the fiber
type InstructionValue struct {
	Key   string
	Value interface{}
}

// KeySearch Search a key in fiber's keys / values array
func KeySearch(name string) *InstructionValue {
	idx := sort.Search(len(FiberValues), func(i int) bool {
		return FiberValues[i].Key == name
	})
	if idx == len(FiberValues) {
		fmt.Println("FIBER ERROR: Unable to find the fiber's var: ", name)
		return &InstructionValue{
			Value: nil,
		}
	} else {
		return FiberValues[idx]
	}
}

// SetValue create or update a key / value in fiber's values
func SetValue(key string, value interface{}) {
	for i := 0; i < len(FiberValues); i++ {
		if FiberValues[i].Key == key {
			FiberValues[i].Value = value
			return
		}
	}
	newValue := &InstructionValue{
		Key:   key,
		Value: value,
	}
	FiberValues = append(FiberValues, newValue)
}
