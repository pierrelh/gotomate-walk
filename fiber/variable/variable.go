package variable

import (
	"fmt"
)

// FiberVariable define the fiber's key / values
var FiberVariable []*InstructionVariable

// InstructionVariable Define a key / value of the fiber
type InstructionVariable struct {
	Key   string
	Value interface{}
}

// SearchVariable Search a key in fiber's keys / values array
func SearchVariable(name string) *InstructionVariable {
	idx := -1
	for i := 0; i < len(FiberVariable); i++ {
		if FiberVariable[i].Key == name {
			idx = i
			break
		}
	}
	if idx != -1 {
		return FiberVariable[idx]
	} else {
		fmt.Println("FIBER ERROR: Unable to find the fiber's var: ", name)
		return &InstructionVariable{
			Value: nil,
		}
	}
}

// SetVariable create or update a key / value in fiber's variable
func SetVariable(key string, value interface{}) {
	for i := 0; i < len(FiberVariable); i++ {
		if FiberVariable[i].Key == key {
			FiberVariable[i].Value = value
			return
		}
	}
	newValue := &InstructionVariable{
		Key:   key,
		Value: value,
	}
	FiberVariable = append(FiberVariable, newValue)
}
