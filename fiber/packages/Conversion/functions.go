package conversion

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"
	"strconv"
)

// BoolToFloat Convert a bool to a float
func BoolToFloat(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting Bool to Float ...")

	var input bool
	if isVar := instructionData.FieldByName("InputIsVar").Interface().(bool); isVar {
		inputVarName := instructionData.FieldByName("InputVarName").Interface().(string)
		if val := variable.SearchVariable(inputVarName).Value; val != nil {
			input = val.(bool)
		} else {
			finished <- true
			return -1
		}
	} else {
		input = instructionData.FieldByName("Input").Interface().(bool)
	}

	var value float64
	if input {
		value = 1
	} else {
		value = 0
	}

	name := instructionData.FieldByName("Output").Interface().(string)
	variable.SetVariable(name, value)
	finished <- true
	return -1
}

// BoolToInt Convert a bool to an int
func BoolToInt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting Bool to Int ...")

	var input bool
	if isVar := instructionData.FieldByName("InputIsVar").Interface().(bool); isVar {
		inputVarName := instructionData.FieldByName("InputVarName").Interface().(string)
		if val := variable.SearchVariable(inputVarName).Value; val != nil {
			input = val.(bool)
		} else {
			finished <- true
			return -1
		}
	} else {
		input = instructionData.FieldByName("Input").Interface().(bool)
	}

	var value int
	if input {
		value = 1
	} else {
		value = 0
	}

	name := instructionData.FieldByName("Output").Interface().(string)
	variable.SetVariable(name, value)
	finished <- true
	return -1
}

// BoolToString Convert a bool to a string
func BoolToString(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting Bool to String ...")

	var input bool
	if isVar := instructionData.FieldByName("InputIsVar").Interface().(bool); isVar {
		inputVarName := instructionData.FieldByName("InputVarName").Interface().(string)
		if val := variable.SearchVariable(inputVarName).Value; val != nil {
			input = val.(bool)
		} else {
			finished <- true
			return -1
		}
	} else {
		input = instructionData.FieldByName("Input").Interface().(bool)
	}

	name := instructionData.FieldByName("Output").Interface().(string)
	value := strconv.FormatBool(input)
	variable.SetVariable(name, value)
	finished <- true
	return -1
}

// FloatToInt Convert a float to an int
func FloatToInt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting Float to Int ...")

	var input float64
	if isVar := instructionData.FieldByName("InputIsVar").Interface().(bool); isVar {
		inputVarName := instructionData.FieldByName("InputVarName").Interface().(string)
		if val := variable.SearchVariable(inputVarName).Value; val != nil {
			input = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		input = instructionData.FieldByName("Input").Interface().(float64)
	}

	name := instructionData.FieldByName("Output").Interface().(string)
	value := int(input)
	variable.SetVariable(name, value)
	finished <- true
	return -1
}

// FloatToString Convert a float to a string
func FloatToString(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting Float to String ...")

	var input float64
	if isVar := instructionData.FieldByName("InputIsVar").Interface().(bool); isVar {
		inputVarName := instructionData.FieldByName("InputVarName").Interface().(string)
		if val := variable.SearchVariable(inputVarName).Value; val != nil {
			input = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		input = instructionData.FieldByName("Input").Interface().(float64)
	}

	name := instructionData.FieldByName("Output").Interface().(string)
	value := strconv.FormatFloat(input, 'E', -1, 64)
	variable.SetVariable(name, value)
	finished <- true
	return -1
}

// IntToFloat Convert an int to a float
func IntToFloat(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting Int to Float ...")

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
	value := float64(input)
	variable.SetVariable(name, value)
	finished <- true
	return -1
}

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

// StringToFloat Convert a string to a float
func StringToFloat(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting String to Float ...")

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
	value, _ := strconv.ParseFloat(input, 64)
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
