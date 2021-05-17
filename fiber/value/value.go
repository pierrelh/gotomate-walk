package value

import (
	"fmt"
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
	idx := -1
	for i := 0; i < len(FiberValues); i++ {
		if FiberValues[i].Key == name {
			idx = i
			break
		}
	}
	if idx != -1 {
		return FiberValues[idx]
	} else {
		fmt.Println("FIBER ERROR: Unable to find the fiber's var: ", name)
		return &InstructionValue{
			Value: nil,
		}
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
