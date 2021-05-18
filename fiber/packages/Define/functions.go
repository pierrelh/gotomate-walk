package define

import (
	"fmt"
	variable "gotomate/fiber/variable"
	"reflect"
)

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
