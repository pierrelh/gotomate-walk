package define

import (
	"fmt"
	variable "gotomate/fiber/variable"
	"reflect"
	"strconv"
	"strings"
)

// ArrayOfBool Define an array of bool in a flow
func ArrayOfBool(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Defining an array of bool ...")

	var value string
	if isVar := instructionData.FieldByName("IsVar").Interface().(bool); isVar {
		varName := instructionData.FieldByName("VarName").Interface().(string)
		if val := variable.SearchVariable(varName).Value; val != nil {
			value = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", varName)
			finished <- true
			return -1
		}
	} else {
		value = instructionData.FieldByName("Value").Interface().(string)
	}

	array := strings.Split(value, ",")
	var boolArray []bool
	for i := 0; i < len(array); i++ {
		boolvalue, _ := strconv.ParseBool(array[i])
		boolArray = append(boolArray, boolvalue)
	}

	name := instructionData.FieldByName("Name").Interface().(string)
	variable.SetVariable(name, boolArray)
	finished <- true
	return -1
}

// ArrayOfInt Define an array of int in a flow
func ArrayOfInt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Defining an array of int ...")

	var value string
	if isVar := instructionData.FieldByName("IsVar").Interface().(bool); isVar {
		varName := instructionData.FieldByName("VarName").Interface().(string)
		if val := variable.SearchVariable(varName).Value; val != nil {
			value = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", varName)
			finished <- true
			return -1
		}
	} else {
		value = instructionData.FieldByName("Value").Interface().(string)
	}

	array := strings.Split(value, ",")
	var intArray []int
	for i := 0; i < len(array); i++ {
		intValue, _ := strconv.Atoi(array[i])
		intArray = append(intArray, intValue)
	}

	name := instructionData.FieldByName("Name").Interface().(string)
	variable.SetVariable(name, intArray)
	finished <- true
	return -1
}

// ArrayOfString Define an array of string in a flow
func ArrayOfString(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Defining an array of string ...")

	var value string
	if isVar := instructionData.FieldByName("IsVar").Interface().(bool); isVar {
		varName := instructionData.FieldByName("VarName").Interface().(string)
		if val := variable.SearchVariable(varName).Value; val != nil {
			value = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", varName)
			finished <- true
			return -1
		}
	} else {
		value = instructionData.FieldByName("Value").Interface().(string)
	}

	stringArray := strings.Split(value, ",")
	name := instructionData.FieldByName("Name").Interface().(string)
	variable.SetVariable(name, stringArray)
	finished <- true
	return -1
}

// Bool Define a bool value in a flow
func Bool(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Defining a bool value ...")
	var value bool

	if isVar := instructionData.FieldByName("IsVar").Interface().(bool); isVar {
		varName := instructionData.FieldByName("VarName").Interface().(string)
		if val := variable.SearchVariable(varName).Value; val != nil {
			value = val.(bool)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", varName)
			finished <- true
			return -1
		}
	} else {
		value = instructionData.FieldByName("Value").Interface().(bool)
	}

	name := instructionData.FieldByName("Name").Interface().(string)
	variable.SetVariable(name, value)
	finished <- true
	return -1
}

// Float Define a float value in a flow
func Float(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Defining a float value ...")
	var value float64

	if isVar := instructionData.FieldByName("IsVar").Interface().(bool); isVar {
		varName := instructionData.FieldByName("VarName").Interface().(string)
		if val := variable.SearchVariable(varName).Value; val != nil {
			value = val.(float64)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", varName)
			finished <- true
			return -1
		}
	} else {
		value = instructionData.FieldByName("Value").Interface().(float64)
	}

	name := instructionData.FieldByName("Name").Interface().(string)
	variable.SetVariable(name, value)
	finished <- true
	return -1
}

// Int Define an int value in a flow
func Int(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Defining an int value ...")
	var value int

	if isVar := instructionData.FieldByName("IsVar").Interface().(bool); isVar {
		varName := instructionData.FieldByName("VarName").Interface().(string)
		if val := variable.SearchVariable(varName).Value; val != nil {
			value = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", varName)
			finished <- true
			return -1
		}
	} else {
		value = instructionData.FieldByName("Value").Interface().(int)
	}

	name := instructionData.FieldByName("Name").Interface().(string)
	variable.SetVariable(name, value)
	finished <- true
	return -1
}

// String Define a string value in a flow
func String(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Defining a string value ...")
	var value string

	if isVar := instructionData.FieldByName("IsVar").Interface().(bool); isVar {
		varName := instructionData.FieldByName("VarName").Interface().(string)
		if val := variable.SearchVariable(varName).Value; val != nil {
			value = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", varName)
			finished <- true
			return -1
		}
	} else {
		value = instructionData.FieldByName("Value").Interface().(string)
	}

	name := instructionData.FieldByName("Name").Interface().(string)
	variable.SetVariable(name, value)
	finished <- true
	return -1
}
