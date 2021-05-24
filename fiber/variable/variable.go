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

// GetVariableType get the type of a value
func GetVariableType(value interface{}) string {
	switch value.(type) {
	case []bool:
		return "[]bool"
	case []complex64:
		return "[]complex64"
	case []complex128:
		return "[]complex128"
	case []float32:
		return "[]float32"
	case []float64:
		return "[]float64"
	case []int:
		return "[]int"
	case []int8:
		return "[]int8"
	case []int16:
		return "[]int16"
	case []int32:
		return "[]int32"
	case []int64:
		return "[]int64"
	case []string:
		return "[]string"
	case []uint:
		return "[]uint"
	case []uint8:
		return "[]uint8"
	case []uint16:
		return "[]uint16"
	case []uint32:
		return "[]uint32"
	case []uint64:
		return "[]uint64"
	case []uintptr:
		return "[]uintptr"
	case bool:
		return "bool"
	case complex64:
		return "complex64"
	case complex128:
		return "complex128"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case int:
		return "int"
	case int8:
		return "int8"
	case int16:
		return "int16"
	case int32:
		return "int32"
	case int64:
		return "int64"
	case string:
		return "string"
	case uint:
		return "uint"
	case uint8:
		return "uint8"
	case uint16:
		return "uint16"
	case uint32:
		return "uint32"
	case uint64:
		return "uint64"
	case uintptr:
		return "uintptr"
	default:
		fmt.Println("FIBER WARNING: Unknown type")
		return ""
	}
}
