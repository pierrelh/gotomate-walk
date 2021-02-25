package fiber

import (
	"reflect"
)

//VarSearch search a var in the fiber & return the instruction
func (fiber *Fiber) VarSearch(name string) *Instruction {
	for i := 0; i < len(fiber.Instructions); i++ {
		instruction := fiber.Instructions[i]
		field := reflect.ValueOf(instruction.Data).Elem().FieldByName("Var")
		if field.IsValid() {
			value := field.Interface().(string)
			if value != "" {
				if value == name {
					return instruction
				}
				continue
			}
		} else {
			continue
		}
	}
	return nil
}
