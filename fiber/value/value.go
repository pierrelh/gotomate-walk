package value

import "reflect"

// FiberValues define the fiber's key / values
var FiberValues []*InstructionValue

// InstructionValue Define a key / value of the fiber
type InstructionValue struct {
	Key   string
	Value interface{}
}

// KeySearch Search a key in fiber's keys / values array
func KeySearch(name string) *InstructionValue {
	for i := 0; i < len(FiberValues); i++ {
		value := FiberValues[i]
		key := reflect.ValueOf(value).Elem().FieldByName("Key").Interface().(string)
		if key != "" {
			if key == name {
				return value
			}
			continue
		}
	}
	return nil
}
