package app

import (
	"reflect"
)

func varSearch(name string) *FiberInstruction {
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
