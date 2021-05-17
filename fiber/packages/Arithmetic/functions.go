package arithmetic

import (
	"fmt"
	"gotomate/fiber/variable"
	"math"
	"reflect"
	"strconv"
)

// Divide make the divide of two numbers
func Divide(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Dividing ...")
	name := instructionData.FieldByName("Output").Interface().(string)
	SOne := instructionData.FieldByName("ValueOne").Interface().(string)
	one, _ := strconv.Atoi(SOne)
	if oneIsVar := instructionData.FieldByName("ValueOneIsVar").Interface().(bool); oneIsVar {
		if val := variable.SearchVariable(SOne).Value; val != nil {
			one = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", SOne)
			finished <- true
			return -1
		}
	}

	STwo := instructionData.FieldByName("ValueTwo").Interface().(string)
	two, _ := strconv.Atoi(STwo)
	if twoIsVar := instructionData.FieldByName("ValueTwoIsVar").Interface().(bool); twoIsVar {
		if val := variable.SearchVariable(STwo).Value; val != nil {
			two = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", STwo)
			finished <- true
			return -1
		}
	}

	result := one / two
	variable.SetVariable(name, result)
	finished <- true
	return -1
}

// Multiply make the multiplication of two numbers
func Multiply(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Multiplying ...")
	name := instructionData.FieldByName("Output").Interface().(string)
	SOne := instructionData.FieldByName("ValueOne").Interface().(string)
	one, _ := strconv.Atoi(SOne)
	if oneIsVar := instructionData.FieldByName("ValueOneIsVar").Interface().(bool); oneIsVar {
		if val := variable.SearchVariable(SOne).Value; val != nil {
			one = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", SOne)
			finished <- true
			return -1
		}
	}

	STwo := instructionData.FieldByName("ValueTwo").Interface().(string)
	two, _ := strconv.Atoi(STwo)
	if twoIsVar := instructionData.FieldByName("ValueTwoIsVar").Interface().(bool); twoIsVar {
		if val := variable.SearchVariable(STwo).Value; val != nil {
			two = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", STwo)
			finished <- true
			return -1
		}
	}

	result := one * two
	variable.SetVariable(name, result)
	finished <- true
	return -1
}

// Pow make the pow of a number
func Pow(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Pow ...")
	name := instructionData.FieldByName("Output").Interface().(string)
	SOne := instructionData.FieldByName("ValueOne").Interface().(string)
	one, _ := strconv.ParseFloat(SOne, 64)
	if oneIsVar := instructionData.FieldByName("ValueOneIsVar").Interface().(bool); oneIsVar {
		if val := variable.SearchVariable(SOne).Value; val != nil {
			one = val.(float64)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", SOne)
			finished <- true
			return -1
		}
	}

	STwo := instructionData.FieldByName("ValueTwo").Interface().(string)
	two, _ := strconv.ParseFloat(STwo, 64)
	if twoIsVar := instructionData.FieldByName("ValueTwoIsVar").Interface().(bool); twoIsVar {
		if val := variable.SearchVariable(STwo).Value; val != nil {
			two = val.(float64)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", STwo)
			finished <- true
			return -1
		}
	}

	result := math.Pow(one, two)
	variable.SetVariable(name, result)
	finished <- true
	return -1
}

// Sqrt make the Sqrt of a number
func Sqrt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Sqrt ...")
	name := instructionData.FieldByName("Output").Interface().(string)
	SValue := instructionData.FieldByName("Value").Interface().(string)
	value, _ := strconv.ParseFloat(SValue, 64)
	if isVar := instructionData.FieldByName("ValueIsVar").Interface().(bool); isVar {
		if val := variable.SearchVariable(SValue).Value; val != nil {
			value = val.(float64)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", SValue)
			finished <- true
			return -1
		}
	}

	result := math.Sqrt(value)
	variable.SetVariable(name, result)
	finished <- true
	return -1
}

// Substract make the substraction of two numbers
func Substract(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Substracting ...")
	name := instructionData.FieldByName("Output").Interface().(string)
	SOne := instructionData.FieldByName("ValueOne").Interface().(string)
	one, _ := strconv.Atoi(SOne)
	if oneIsVar := instructionData.FieldByName("ValueOneIsVar").Interface().(bool); oneIsVar {
		if val := variable.SearchVariable(SOne).Value; val != nil {
			one = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", SOne)
			finished <- true
			return -1
		}
	}

	STwo := instructionData.FieldByName("ValueTwo").Interface().(string)
	two, _ := strconv.Atoi(STwo)
	if twoIsVar := instructionData.FieldByName("ValueTwoIsVar").Interface().(bool); twoIsVar {
		if val := variable.SearchVariable(STwo).Value; val != nil {
			two = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", STwo)
			finished <- true
			return -1
		}
	}

	result := one - two
	variable.SetVariable(name, result)
	finished <- true
	return -1
}

// Sum make the sum of two numbers
func Sum(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Additioning ...")
	name := instructionData.FieldByName("Output").Interface().(string)
	SOne := instructionData.FieldByName("ValueOne").Interface().(string)
	one, _ := strconv.Atoi(SOne)
	if oneIsVar := instructionData.FieldByName("ValueOneIsVar").Interface().(bool); oneIsVar {
		if val := variable.SearchVariable(SOne).Value; val != nil {
			one = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", SOne)
			finished <- true
			return -1
		}
	}

	STwo := instructionData.FieldByName("ValueTwo").Interface().(string)
	two, _ := strconv.Atoi(STwo)
	if twoIsVar := instructionData.FieldByName("ValueTwoIsVar").Interface().(bool); twoIsVar {
		if val := variable.SearchVariable(STwo).Value; val != nil {
			two = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", STwo)
			finished <- true
			return -1
		}
	}

	result := one + two
	variable.SetVariable(name, result)
	finished <- true
	return -1
}
