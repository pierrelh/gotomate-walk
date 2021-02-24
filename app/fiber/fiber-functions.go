package fiber

import (
	"reflect"
)

//VarSearch search a var in the fiber & return the instruction
func (fiber *Fiber) VarSearch(name string) *Instruction {
	for i := 0; i < len(fiber.Instructions); i++ {
		instruction := fiber.Instructions[i]
		val := reflect.ValueOf(instruction.Data).Elem()
		err := val.FieldByName("Var").Interface().(string)
		if err != "" {
			if err == name {
				return instruction
			}
			continue
		}
	}
	return nil
}
