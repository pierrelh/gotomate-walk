package algorithmic

import (
	"fmt"
	variable "gotomate/fiber/variable"
	"reflect"
)

// For Execute a for loop
func For(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: For Statement ...")

	var valueOne int
	varOneName := instructionData.FieldByName("VarOneName").Interface().(string)
	if val := variable.SearchVariable(varOneName).Value; val != nil {
		valueOne = val.(int)
	} else {
		finished <- true
		return -1
	}

	var valueTwo int
	if twoIsVar := instructionData.FieldByName("TwoIsVar").Interface().(bool); twoIsVar {
		varTwoName := instructionData.FieldByName("VarTwoName").Interface().(string)
		if val := variable.SearchVariable(varTwoName).Value; val != nil {
			valueTwo = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		valueTwo = instructionData.FieldByName("ValueTwo").Interface().(int)
	}

	comparator := instructionData.FieldByName("Comparator").Interface().(string)
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
		if valueOne > valueTwo {
			statement = true
		}
	case ">=":
		if valueOne >= valueTwo {
			statement = true
		}
	case "<":
		if valueOne < valueTwo {
			statement = true
		}
	case "<=":
		if valueOne <= valueTwo {
			statement = true
		}
	}

	if statement {
		var increment int
		if incrementIsVar := instructionData.FieldByName("IncrementIsVar").Interface().(bool); incrementIsVar {
			incrementVarName := instructionData.FieldByName("IncrementVarName").Interface().(string)
			if val := variable.SearchVariable(incrementVarName).Value; val != nil {
				increment = val.(int)
			} else {
				finished <- true
				return -1
			}
		} else {
			increment = instructionData.FieldByName("Increment").Interface().(int)
		}

		valueOne = valueOne + increment
		variable.SetVariable(varOneName, valueOne)
		finished <- true
		return -1
	} else {
		var falseInstruction int
		if falseInstructionIsVar := instructionData.FieldByName("FalseInstructionIsVar").Interface().(bool); falseInstructionIsVar {
			falseInstructionVarName := instructionData.FieldByName("FalseInstructionVarName").Interface().(string)
			if val := variable.SearchVariable(falseInstructionVarName).Value; val != nil {
				falseInstruction = val.(int)
			} else {
				finished <- true
				return -1
			}
		} else {
			falseInstruction = instructionData.FieldByName("FalseInstruction").Interface().(int)
		}
		finished <- true
		return falseInstruction
	}
}

// If Compare if a statement is true
func If(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: If Statement ...")
	var valueOne interface{}

	if oneIsVar := instructionData.FieldByName("OneIsVar").Interface().(bool); oneIsVar {
		varOneName := instructionData.FieldByName("VarOneName").Interface().(string)
		if val := variable.SearchVariable(varOneName).Value; val != nil {
			valueOne = val
		} else {
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
			finished <- true
			return -1
		}
	} else {
		valueTwo = instructionData.FieldByName("ValueTwo").Interface()
	}

	comparator := instructionData.FieldByName("Comparator").Interface().(string)
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
		var falseInstruction int
		if falseInstructionIsVar := instructionData.FieldByName("FalseInstructionIsVar").Interface().(bool); falseInstructionIsVar {
			falseInstructionVarName := instructionData.FieldByName("FalseInstructionVarName").Interface().(string)
			if val := variable.SearchVariable(falseInstructionVarName).Value; val != nil {
				falseInstruction = val.(int)
			} else {
				finished <- true
				return -1
			}
		} else {
			falseInstruction = instructionData.FieldByName("FalseInstruction").Interface().(int)
		}
		return falseInstruction
	}
}
