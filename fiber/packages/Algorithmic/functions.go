package algorithmic

import (
	"fmt"
	variable "gotomate/fiber/variable"
	"reflect"
)

// If Compare if a statement is true
func If(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: If Statement ...")
	var valueOne interface{}

	if oneIsVar := instructionData.FieldByName("OneIsVar").Interface().(bool); oneIsVar {
		varOneName := instructionData.FieldByName("VarOneName").Interface().(string)
		if val := variable.SearchVariable(varOneName).Value; val != nil {
			valueOne = val
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", varOneName)
			finished <- true
			return -1
		}
	} else {
		valueOne = instructionData.FieldByName("ValueOne").Interface()
	}

	var valueTwo interface{}

	if twoIsVar := instructionData.FieldByName("TwoIsVar").Interface().(bool); twoIsVar {
		varTwoName := instructionData.FieldByName("VarTwoName").Interface().(string)
		if val := variable.SearchVariable(varTwoName).Value; val != nil {
			valueTwo = val
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", varTwoName)
			finished <- true
			return -1
		}
	} else {
		valueTwo = instructionData.FieldByName("ValueTwo").Interface()
	}
	
	comparator := instructionData.FieldByName("Comparator").Interface().(string)
	falseInstruction := instructionData.FieldByName("FalseInstruction").Interface().(int)
	statement := false
	switch comparator {
	case "==":
		if valueOne == valueTwo {
			statement = true
		}
	case "!=":
		if valueOne != valueTwo {
			statement = true
		}
	case ">":
		if valueOne.(int) > valueTwo.(int) {
			statement = true
		}
	case ">=":
		if valueOne.(int) >= valueTwo.(int) {
			statement = true
		}
	case "<":
		if valueOne.(int) < valueTwo.(int) {
			statement = true
		}
	case "<=":
		if valueOne.(int) <= valueTwo.(int) {
			statement = true
		}
	}
	finished <- true
	if statement {
		return -1
	} else {
		return falseInstruction
	}
}
