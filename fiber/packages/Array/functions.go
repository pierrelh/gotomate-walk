package array

import (
	"fmt"
	"gotomate/fiber/variable"
	"math/rand"
	"reflect"
	"time"
)

// GetArrayLength Get the current length of an array
func GetArrayLength(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Get the lenth of an array ...")

	var bools []bool
	var ints []int
	var strings []string
	var typeOfArray string
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		switch val := val.(type) {
		case []bool:
			bools = val
			typeOfArray = "bool"
		case []int:
			ints = val
			typeOfArray = "int"
		case []string:
			strings = val
			typeOfArray = "string"
		default:
			fmt.Println("FIBER WARNING: Type error on length array")
			finished <- true
			return -1
		}
	} else {
		finished <- true
		return -1
	}

	name := instructionData.FieldByName("Output").Interface().(string)

	if typeOfArray == "bool" {
		variable.SetVariable(name, len(bools))
	} else if typeOfArray == "int" {
		variable.SetVariable(name, len(ints))
	} else {
		variable.SetVariable(name, len(strings))
	}
	finished <- true
	return -1
}

// PopAt Pop the wanted index of an array
func PopAt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Pop array at index ...")

	var bools []bool
	var ints []int
	var strings []string
	var typeOfArray string
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		switch val := val.(type) {
		case []bool:
			bools = val
			typeOfArray = "bool"
		case []int:
			ints = val
			typeOfArray = "int"
		case []string:
			strings = val
			typeOfArray = "string"
		default:
			fmt.Println("FIBER WARNING: Type error on popping array")
			finished <- true
			return -1
		}
	} else {
		finished <- true
		return -1
	}

	var index int
	if indexIsVar := instructionData.FieldByName("IndexIsVar").Interface().(bool); indexIsVar {
		indexVarName := instructionData.FieldByName("IndexVarName").Interface().(string)
		if val := variable.SearchVariable(indexVarName).Value; val != nil {
			index = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		index = instructionData.FieldByName("Index").Interface().(int)
	}

	name := instructionData.FieldByName("Output").Interface().(string)

	if typeOfArray == "bool" {
		popped := bools[index]
		copy(bools[index:], bools[index+1:])
		bools = bools[:len(bools)-1]
		variable.SetVariable(name, popped)
		variable.SetVariable(arrayVarName, bools)
	} else if typeOfArray == "int" {
		popped := ints[index]
		copy(ints[index:], ints[index+1:])
		ints = ints[:len(ints)-1]
		variable.SetVariable(name, popped)
		variable.SetVariable(arrayVarName, ints)
	} else {
		popped := strings[index]
		copy(strings[index:], strings[index+1:])
		strings = strings[:len(strings)-1]
		variable.SetVariable(name, popped)
		variable.SetVariable(arrayVarName, strings)
	}
	finished <- true
	return -1
}

// PopLast Pop the last index of an array
func PopLast(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Pop array at end ...")

	var bools []bool
	var ints []int
	var strings []string
	var typeOfArray string
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		switch val := val.(type) {
		case []bool:
			bools = val
			typeOfArray = "bool"
		case []int:
			ints = val
			typeOfArray = "int"
		case []string:
			strings = val
			typeOfArray = "string"
		default:
			fmt.Println("FIBER WARNING: Type error on popping array")
			finished <- true
			return -1
		}
	} else {
		finished <- true
		return -1
	}

	name := instructionData.FieldByName("Output").Interface().(string)

	if typeOfArray == "bool" {
		popped := bools[len(bools)-1]
		bools = bools[:len(bools)-1]
		variable.SetVariable(name, popped)
		variable.SetVariable(arrayVarName, bools)
	} else if typeOfArray == "int" {
		popped := ints[len(ints)-1]
		ints = ints[:len(ints)-1]
		variable.SetVariable(name, popped)
		variable.SetVariable(arrayVarName, ints)
	} else {
		popped := strings[len(strings)-1]
		strings = strings[:len(strings)-1]
		variable.SetVariable(name, popped)
		variable.SetVariable(arrayVarName, strings)
	}
	finished <- true
	return -1
}

// PushAt Push a value at the wanted index of an array
func PushAt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Push value in array at index ...")

	var bools []bool
	var ints []int
	var strings []string
	var typeOfArray string
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		switch val := val.(type) {
		case []bool:
			bools = val
			typeOfArray = "bool"
		case []int:
			ints = val
			typeOfArray = "int"
		case []string:
			strings = val
			typeOfArray = "string"
		default:
			fmt.Println("FIBER WARNING: Type error on pushing value in array")
			finished <- true
			return -1
		}
	} else {
		finished <- true
		return -1
	}

	var index int
	if indexIsVar := instructionData.FieldByName("IndexIsVar").Interface().(bool); indexIsVar {
		indexVarName := instructionData.FieldByName("IndexVarName").Interface().(string)
		if val := variable.SearchVariable(indexVarName).Value; val != nil {
			index = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		index = instructionData.FieldByName("Index").Interface().(int)
	}

	var boolValue bool
	var intValue int
	var stringValue string
	valueVarName := instructionData.FieldByName("ValueVarName").Interface().(string)
	if val := variable.SearchVariable(valueVarName).Value; val != nil {
		switch val := val.(type) {
		case bool:
			boolValue = val
		case int:
			intValue = int(val)
		case string:
			stringValue = val
		default:
			fmt.Println("FIBER WARNING: Type error on pushing value in array")
			finished <- true
			return -1
		}
	} else {
		finished <- true
		return -1
	}

	if typeOfArray == "bool" {
		bools = append(bools, false)
		copy(bools[index+1:], bools[index:])
		bools[index] = boolValue
		variable.SetVariable(arrayVarName, bools)
	} else if typeOfArray == "int" {
		ints = append(ints, 0)
		copy(ints[index+1:], ints[index:])
		ints[index] = intValue
		variable.SetVariable(arrayVarName, ints)
	} else {
		strings = append(strings, "")
		copy(strings[index+1:], strings[index:])
		strings[index] = stringValue
		variable.SetVariable(arrayVarName, strings)
	}
	finished <- true
	return -1
}

// PushLast Push a value at the end of an array
func PushLast(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Push value in array at end ...")

	var bools []bool
	var ints []int
	var strings []string
	var typeOfArray string
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		switch val := val.(type) {
		case []bool:
			bools = val
			typeOfArray = "bool"
		case []int:
			ints = val
			typeOfArray = "int"
		case []string:
			strings = val
			typeOfArray = "string"
		default:
			fmt.Println("FIBER WARNING: Type error on pushing value in array")
			finished <- true
			return -1
		}
	} else {
		finished <- true
		return -1
	}

	var boolValue bool
	var intValue int
	var stringValue string
	valueVarName := instructionData.FieldByName("ValueVarName").Interface().(string)
	if val := variable.SearchVariable(valueVarName).Value; val != nil {
		switch val := val.(type) {
		case bool:
			boolValue = val
		case int:
			intValue = val
		case string:
			stringValue = val
		default:
			fmt.Println("FIBER WARNING: Type error on pushing value in array")
			finished <- true
			return -1
		}
	} else {
		finished <- true
		return -1
	}

	if typeOfArray == "bool" {
		bools = append(bools, boolValue)
		variable.SetVariable(arrayVarName, bools)
	} else if typeOfArray == "int" {
		ints = append(ints, intValue)
		variable.SetVariable(arrayVarName, ints)
	} else {
		strings = append(strings, stringValue)
		variable.SetVariable(arrayVarName, strings)
	}
	finished <- true
	return -1
}

// RemoveAt Remove a value by the index, of an array
func RemoveAt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Remove value from array at index ...")

	var bools []bool
	var ints []int
	var strings []string
	var typeOfArray string
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		switch val := val.(type) {
		case []bool:
			bools = val
			typeOfArray = "bool"
		case []int:
			ints = val
			typeOfArray = "int"
		case []string:
			strings = val
			typeOfArray = "string"
		default:
			fmt.Println("FIBER WARNING: Type error on removing value in array")
			finished <- true
			return -1
		}
	} else {
		finished <- true
		return -1
	}

	var index int
	if indexIsVar := instructionData.FieldByName("IndexIsVar").Interface().(bool); indexIsVar {
		indexVarName := instructionData.FieldByName("IndexVarName").Interface().(string)
		if val := variable.SearchVariable(indexVarName).Value; val != nil {
			index = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		index = instructionData.FieldByName("Index").Interface().(int)
	}

	if typeOfArray == "bool" {
		copy(bools[index:], bools[index+1:])
		bools = bools[:len(bools)-1]
		variable.SetVariable(arrayVarName, bools)
	} else if typeOfArray == "int" {
		copy(ints[index:], ints[index+1:])
		ints = ints[:len(ints)-1]
		variable.SetVariable(arrayVarName, ints)
	} else {
		copy(strings[index:], strings[index+1:])
		strings = strings[:len(strings)-1]
		variable.SetVariable(arrayVarName, strings)
	}
	finished <- true
	return -1
}

// RemoveLast Remove the last value of an array
func RemoveLast(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Remove value from array at end ...")

	var bools []bool
	var ints []int
	var strings []string
	var typeOfArray string
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		switch val := val.(type) {
		case []bool:
			bools = val
			typeOfArray = "bool"
		case []int:
			ints = val
			typeOfArray = "int"
		case []string:
			strings = val
			typeOfArray = "string"
		default:
			fmt.Println("FIBER WARNING: Type error on removing value in array")
			finished <- true
			return -1
		}
	} else {
		finished <- true
		return -1
	}

	if typeOfArray == "bool" {
		bools = bools[:len(bools)-1]
		variable.SetVariable(arrayVarName, bools)
	} else if typeOfArray == "int" {
		ints = ints[:len(ints)-1]
		variable.SetVariable(arrayVarName, ints)
	} else {
		strings = strings[:len(strings)-1]
		variable.SetVariable(arrayVarName, strings)
	}
	finished <- true
	return -1
}

// Shuffle an array
func Shuffle(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Shuffling array ...")

	var bools []bool
	var ints []int
	var strings []string
	var typeOfArray string
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		switch val := val.(type) {
		case []bool:
			bools = val
			typeOfArray = "bool"
		case []int:
			ints = val
			typeOfArray = "int"
		case []string:
			strings = val
			typeOfArray = "string"
		default:
			fmt.Println("FIBER WARNING: Type error on updating value in array")
			finished <- true
			return -1
		}
	} else {
		finished <- true
		return -1
	}

	rand.Seed(time.Now().UnixNano())

	if typeOfArray == "bool" {
		rand.Shuffle(len(bools), func(i, j int) { bools[i], bools[j] = bools[j], bools[i] })
		variable.SetVariable(arrayVarName, bools)
	} else if typeOfArray == "int" {
		rand.Shuffle(len(ints), func(i, j int) { ints[i], ints[j] = ints[j], ints[i] })
		variable.SetVariable(arrayVarName, ints)
	} else {
		rand.Shuffle(len(strings), func(i, j int) { strings[i], strings[j] = strings[j], strings[i] })
		variable.SetVariable(arrayVarName, strings)
	}
	finished <- true
	return -1
}

// UpdateValue Update a value of an array by index
func UpdateValue(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Updating value in array by index ...")

	var bools []bool
	var ints []int
	var strings []string
	var typeOfArray string
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		switch val := val.(type) {
		case []bool:
			bools = val
			typeOfArray = "bool"
		case []int:
			ints = val
			typeOfArray = "int"
		case []string:
			strings = val
			typeOfArray = "string"
		default:
			fmt.Println("FIBER WARNING: Type error on updating value in array")
			finished <- true
			return -1
		}
	} else {
		finished <- true
		return -1
	}

	var index int
	if indexIsVar := instructionData.FieldByName("IndexIsVar").Interface().(bool); indexIsVar {
		indexVarName := instructionData.FieldByName("IndexVarName").Interface().(string)
		if val := variable.SearchVariable(indexVarName).Value; val != nil {
			index = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		index = instructionData.FieldByName("Index").Interface().(int)
	}

	var boolValue bool
	var intValue int
	var stringValue string
	valueVarName := instructionData.FieldByName("ValueVarName").Interface().(string)
	if val := variable.SearchVariable(valueVarName).Value; val != nil {
		switch val := val.(type) {
		case bool:
			boolValue = val
		case int:
			intValue = int(val)
		case string:
			stringValue = val
		default:
			fmt.Println("FIBER WARNING: Type error on pushing value in array")
			finished <- true
			return -1
		}
	} else {
		finished <- true
		return -1
	}

	if typeOfArray == "bool" {
		bools[index] = boolValue
		variable.SetVariable(arrayVarName, bools)
	} else if typeOfArray == "int" {
		ints[index] = intValue
		variable.SetVariable(arrayVarName, ints)
	} else {
		strings[index] = stringValue
		variable.SetVariable(arrayVarName, strings)
	}
	finished <- true
	return -1
}
