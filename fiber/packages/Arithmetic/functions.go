package arithmetic

import (
	"fmt"
	"gotomate/fiber/variable"
	"math"
	"reflect"
)

// Divide make the divide of two numbers
func Divide(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Dividing ...")

	var one float64
	if oneIsVar := instructionData.FieldByName("ValueOneIsVar").Interface().(bool); oneIsVar {
		varOneName := instructionData.FieldByName("VarOneName").Interface().(string)
		if val := variable.SearchVariable(varOneName).Value; val != nil {
			one = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		one = instructionData.FieldByName("ValueOne").Interface().(float64)
	}

	var two float64
	if twoIsVar := instructionData.FieldByName("ValueTwoIsVar").Interface().(bool); twoIsVar {
		varTwoName := instructionData.FieldByName("VarTwoName").Interface().(string)
		if val := variable.SearchVariable(varTwoName).Value; val != nil {
			two = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		two = instructionData.FieldByName("ValueTwo").Interface().(float64)
	}

	result := one / two
	name := instructionData.FieldByName("Output").Interface().(string)
	variable.SetVariable(name, result)
	finished <- true
	return -1
}

// Multiply make the multiplication of two numbers
func Multiply(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Multiplying ...")
	var one float64

	if oneIsVar := instructionData.FieldByName("ValueOneIsVar").Interface().(bool); oneIsVar {
		varOneName := instructionData.FieldByName("VarOneName").Interface().(string)
		if val := variable.SearchVariable(varOneName).Value; val != nil {
			one = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		one = instructionData.FieldByName("ValueOne").Interface().(float64)
	}

	var two float64
	if twoIsVar := instructionData.FieldByName("ValueTwoIsVar").Interface().(bool); twoIsVar {
		varTwoName := instructionData.FieldByName("VarTwoName").Interface().(string)
		if val := variable.SearchVariable(varTwoName).Value; val != nil {
			two = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		two = instructionData.FieldByName("ValueTwo").Interface().(float64)
	}

	result := one * two
	name := instructionData.FieldByName("Output").Interface().(string)
	variable.SetVariable(name, result)
	finished <- true
	return -1
}

// Pow make the pow of a number
func Pow(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Pow ...")

	var one float64
	if oneIsVar := instructionData.FieldByName("ValueOneIsVar").Interface().(bool); oneIsVar {
		varOneName := instructionData.FieldByName("VarOneName").Interface().(string)
		if val := variable.SearchVariable(varOneName).Value; val != nil {
			one = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		one = instructionData.FieldByName("ValueOne").Interface().(float64)
	}

	var two float64
	if twoIsVar := instructionData.FieldByName("ValueTwoIsVar").Interface().(bool); twoIsVar {
		varTwoName := instructionData.FieldByName("VarTwoName").Interface().(string)
		if val := variable.SearchVariable(varTwoName).Value; val != nil {
			two = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		two = instructionData.FieldByName("ValueTwo").Interface().(float64)
	}

	result := math.Pow(one, two)
	name := instructionData.FieldByName("Output").Interface().(string)
	variable.SetVariable(name, result)
	finished <- true
	return -1
}

// Sqrt make the Sqrt of a number
func Sqrt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Sqrt ...")

	var value float64
	if isVar := instructionData.FieldByName("ValueIsVar").Interface().(bool); isVar {
		varName := instructionData.FieldByName("VarName").Interface().(string)
		if val := variable.SearchVariable(varName).Value; val != nil {
			value = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		value = instructionData.FieldByName("Value").Interface().(float64)
	}

	result := math.Sqrt(value)
	name := instructionData.FieldByName("Output").Interface().(string)
	variable.SetVariable(name, result)
	finished <- true
	return -1
}

// Substract make the substraction of two numbers
func Substract(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Substracting ...")

	var one float64
	if oneIsVar := instructionData.FieldByName("ValueOneIsVar").Interface().(bool); oneIsVar {
		varOneName := instructionData.FieldByName("VarOneName").Interface().(string)
		if val := variable.SearchVariable(varOneName).Value; val != nil {
			one = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		one = instructionData.FieldByName("ValueOne").Interface().(float64)
	}

	var two float64
	if twoIsVar := instructionData.FieldByName("ValueTwoIsVar").Interface().(bool); twoIsVar {
		varTwoName := instructionData.FieldByName("VarTwoName").Interface().(string)
		if val := variable.SearchVariable(varTwoName).Value; val != nil {
			two = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		two = instructionData.FieldByName("ValueTwo").Interface().(float64)
	}

	result := one - two
	name := instructionData.FieldByName("Output").Interface().(string)
	variable.SetVariable(name, result)
	finished <- true
	return -1
}

// Sum make the sum of two numbers
func Sum(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Additioning ...")

	var one float64
	if oneIsVar := instructionData.FieldByName("ValueOneIsVar").Interface().(bool); oneIsVar {
		varOneName := instructionData.FieldByName("VarOneName").Interface().(string)
		if val := variable.SearchVariable(varOneName).Value; val != nil {
			one = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		one = instructionData.FieldByName("ValueOne").Interface().(float64)
	}

	var two float64
	if twoIsVar := instructionData.FieldByName("ValueTwoIsVar").Interface().(bool); twoIsVar {
		varTwoName := instructionData.FieldByName("VarTwoName").Interface().(string)
		if val := variable.SearchVariable(varTwoName).Value; val != nil {
			two = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		two = instructionData.FieldByName("ValueTwo").Interface().(float64)
	}

	result := one + two
	name := instructionData.FieldByName("Output").Interface().(string)
	variable.SetVariable(name, result)
	finished <- true
	return -1
}
