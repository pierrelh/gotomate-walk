package conversion

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"
	"strconv"
)

// IntToString Convert an int to a string
func IntToString(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting Int to String ...")

	var input int
	if isVar := instructionData.FieldByName("InputIsVar").Interface().(bool); isVar {
		inputVarName := instructionData.FieldByName("InputVarName").Interface().(string)
		if val := variable.SearchVariable(inputVarName).Value; val != nil {
			input = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		input = instructionData.FieldByName("Input").Interface().(int)
	}

	name := instructionData.FieldByName("Output").Interface().(string)
	value := strconv.Itoa(input)
	variable.SetVariable(name, value)
	finished <- true
	return -1
}

// StringToBool Convert a string to a bool
func StringToBool(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting String to Bool ...")

	var input string
	if isVar := instructionData.FieldByName("InputIsVar").Interface().(bool); isVar {
		inputVarName := instructionData.FieldByName("InputVarName").Interface().(string)
		if val := variable.SearchVariable(inputVarName).Value; val != nil {
			input = val.(string)
		} else {
			finished <- true
			return -1
		}
	} else {
		input = instructionData.FieldByName("Input").Interface().(string)
	}

	name := instructionData.FieldByName("Output").Interface().(string)
	value, _ := strconv.ParseBool(input)
	variable.SetVariable(name, value)
	finished <- true
	return -1
}

// StringToInt Convert a string to an int
func StringToInt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting String to Int ...")

	var input string
	if isVar := instructionData.FieldByName("InputIsVar").Interface().(bool); isVar {
		inputVarName := instructionData.FieldByName("InputVarName").Interface().(string)
		if val := variable.SearchVariable(inputVarName).Value; val != nil {
			input = val.(string)
		} else {
			finished <- true
			return -1
		}
	} else {
		input = instructionData.FieldByName("Input").Interface().(string)
	}

	name := instructionData.FieldByName("Output").Interface().(string)
	value, _ := strconv.Atoi(input)
	variable.SetVariable(name, value)
	finished <- true
	return -1
}
