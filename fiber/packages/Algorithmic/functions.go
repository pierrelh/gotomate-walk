package algorithmic

import (
	"fmt"
	variable "gotomate/fiber/variable"
	"reflect"
)

// For Execute a for loop
func For(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: For Statement ...")

	valueOne, err := variable.GetValue(instructionData, "VarOneName")
	if err != nil {
		finished <- true
		return -1
	}

	valueTwo, err := variable.GetValue(instructionData, "VarTwoName", "TwoIsVar", "ValueTwo")
	if err != nil {
		finished <- true
		return -1
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
		if variable.GetFloat(valueOne) > variable.GetFloat(valueTwo) {
			statement = true
		}
	case ">=":
		if variable.GetFloat(valueOne) >= variable.GetFloat(valueTwo) {
			statement = true
		}
	case "<":
		if variable.GetFloat(valueOne) < variable.GetFloat(valueTwo) {
			statement = true
		}
	case "<=":
		if variable.GetFloat(valueOne) <= variable.GetFloat(valueTwo) {
			statement = true
		}
	}

	if statement {
		increment, err := variable.GetValue(instructionData, "IncrementVarName", "IncrementIsVar", "Increment")
		if err != nil {
			finished <- true
			return -1
		}

		variable.SetVariable(instructionData.FieldByName("VarOneName").Interface().(string), variable.GetFloat(valueOne)+float64(increment.(int)))
		finished <- true
		return -1
	} else {
		falseInstruction, err := variable.GetValue(instructionData, "FalseInstructionVarName", "FalseInstructionIsVar", "FalseInstruction")
		if err != nil {
			finished <- true
			return -1
		}
		finished <- true
		return falseInstruction.(int)
	}
}

// If Compare if a statement is true
func If(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: If Statement ...")

	valueOne, err := variable.GetValue(instructionData, "VarOneName", "OneIsVar", "ValueOne")
	if err != nil {
		finished <- true
		return -1
	}

	valueTwo, err := variable.GetValue(instructionData, "VarTwoName", "TwoIsVar", "ValueTwo")
	if err != nil {
		finished <- true
		return -1
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
		if variable.GetFloat(valueOne) > variable.GetFloat(valueTwo) {
			statement = true
		}
	case ">=":
		if variable.GetFloat(valueOne) >= variable.GetFloat(valueTwo) {
			statement = true
		}
	case "<":
		if variable.GetFloat(valueOne) < variable.GetFloat(valueTwo) {
			statement = true
		}
	case "<=":
		if variable.GetFloat(valueOne) <= variable.GetFloat(valueTwo) {
			statement = true
		}
	}
	finished <- true
	if statement {
		return -1
	} else {
		falseInstruction, err := variable.GetValue(instructionData, "FalseInstructionVarName", "FalseInstructionIsVar", "FalseInstruction")
		if err != nil {
			finished <- true
			return -1
		}
		return falseInstruction.(int)
	}
}
