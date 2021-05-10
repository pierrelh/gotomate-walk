package algorithmic

import (
	"fmt"
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	switch funcName {
	case "DefineInt":
		name := instructionData.FieldByName("Name").Interface().(string)
		value := instructionData.FieldByName("Value").Interface().(int)
		go DefineInt(name, value, finished)
		<-finished
	case "DefineString":
		name := instructionData.FieldByName("Name").Interface().(string)
		value := instructionData.FieldByName("Value").Interface().(string)
		go DefineString(name, value, finished)
		<-finished
	case "DefineBool":
		name := instructionData.FieldByName("Name").Interface().(string)
		value := instructionData.FieldByName("Value").Interface().(bool)
		go DefineBool(name, value, finished)
		<-finished
	case "If":
		isTrue := false
		valueOne := instructionData.FieldByName("ValueOne").Interface()
		valueTwo := instructionData.FieldByName("ValueTwo").Interface()
		comparator := instructionData.FieldByName("Comparator").Interface().(string)
		falseInstruction := instructionData.FieldByName("FalseInstruction").Interface().(int)
		go func() {
			isTrue = If(valueOne, valueTwo, comparator, finished)
		}()
		<-finished
		if !isTrue {
			return falseInstruction
		}
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return -1
}
