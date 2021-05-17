package arithmetic

import (
	"fmt"
	"gotomate/fiber/value"
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
		if val := value.KeySearch(SOne).Value; val != nil {
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
		if val := value.KeySearch(STwo).Value; val != nil {
			two = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", STwo)
			finished <- true
			return -1
		}
	}

	result := one / two
	value.SetValue(name, result)
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
		if val := value.KeySearch(SOne).Value; val != nil {
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
		if val := value.KeySearch(STwo).Value; val != nil {
			two = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", STwo)
			finished <- true
			return -1
		}
	}

	result := one * two
	value.SetValue(name, result)
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
		if val := value.KeySearch(SOne).Value; val != nil {
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
		if val := value.KeySearch(STwo).Value; val != nil {
			two = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", STwo)
			finished <- true
			return -1
		}
	}

	result := one - two
	value.SetValue(name, result)
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
		if val := value.KeySearch(SOne).Value; val != nil {
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
		if val := value.KeySearch(STwo).Value; val != nil {
			two = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", STwo)
			finished <- true
			return -1
		}
	}

	result := one + two
	value.SetValue(name, result)
	finished <- true
	return -1
}
