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

	var array interface{}
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		array = val
	} else {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), len(array.([]bool)))
	case "[]int":
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), len(array.([]int)))
	case "[]string":
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), len(array.([]string)))
	}
	finished <- true
	return -1
}

// GetValue Get a value from an array by index
func GetValue(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting a value from array ...")

	var array interface{}
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		array = val
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

	switch variable.GetVariableType(array) {
	case "[]bool":
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), array.([]bool)[index])
	case "[]int":
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), array.([]int)[index])
	case "[]string":
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), array.([]string)[index])
	}
	finished <- true
	return -1
}

// PopAt Pop the wanted index of an array
func PopAt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Pop array at index ...")

	var array interface{}
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		array = val
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

	switch variable.GetVariableType(array) {
	case "[]bool":
		popped := array.([]bool)[index]
		copy(array.([]bool)[index:], array.([]bool)[index+1:])
		array = array.([]bool)[:len(array.([]bool))-1]
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), popped)
		variable.SetVariable(arrayVarName, array.([]bool))
	case "[]int":
		popped := array.([]int)[index]
		copy(array.([]int)[index:], array.([]int)[index+1:])
		array = array.([]int)[:len(array.([]int))-1]
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), popped)
		variable.SetVariable(arrayVarName, array.([]int))
	case "[]string":
		popped := array.([]string)[index]
		copy(array.([]string)[index:], array.([]string)[index+1:])
		array = array.([]string)[:len(array.([]string))-1]
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), popped)
		variable.SetVariable(arrayVarName, array.([]string))
	}
	finished <- true
	return -1
}

// PopLast Pop the last index of an array
func PopLast(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Pop array at end ...")

	var array interface{}
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		array = val
	} else {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		popped := array.([]bool)[len(array.([]bool))-1]
		array = array.([]bool)[:len(array.([]bool))-1]
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), popped)
		variable.SetVariable(arrayVarName, array.([]bool))
	case "[]int":
		popped := array.([]int)[len(array.([]int))-1]
		array = array.([]int)[:len(array.([]int))-1]
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), popped)
		variable.SetVariable(arrayVarName, array.([]int))
	case "[]string":
		popped := array.([]string)[len(array.([]string))-1]
		array = array.([]string)[:len(array.([]string))-1]
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), popped)
		variable.SetVariable(arrayVarName, array.([]string))
	}
	finished <- true
	return -1
}

// PushAt Push a value at the wanted index of an array
func PushAt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Push value in array at index ...")

	var array interface{}
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		array = val
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

	var value interface{}
	valueVarName := instructionData.FieldByName("ValueVarName").Interface().(string)
	if val := variable.SearchVariable(valueVarName).Value; val != nil {
		value = val
	} else {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		array = append(array.([]bool), false)
		copy(array.([]bool)[index+1:], array.([]bool)[index:])
		array.([]bool)[index] = value.(bool)
		variable.SetVariable(arrayVarName, array.([]bool))
	case "[]int":
		array = append(array.([]int), 0)
		copy(array.([]int)[index+1:], array.([]int)[index:])
		array.([]int)[index] = value.(int)
		variable.SetVariable(arrayVarName, array.([]int))
	case "[]string":
		array = append(array.([]string), "")
		copy(array.([]string)[index+1:], array.([]string)[index:])
		array.([]string)[index] = value.(string)
		variable.SetVariable(arrayVarName, array.([]string))
	}
	finished <- true
	return -1
}

// PushLast Push a value at the end of an array
func PushLast(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Push value in array at end ...")

	var array interface{}
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		array = val
	} else {
		finished <- true
		return -1
	}

	var value interface{}
	valueVarName := instructionData.FieldByName("ValueVarName").Interface().(string)
	if val := variable.SearchVariable(valueVarName).Value; val != nil {
		value = val
	} else {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		array = append(array.([]bool), value.(bool))
		variable.SetVariable(arrayVarName, array.([]bool))
	case "[]int":
		array = append(array.([]int), value.(int))
		variable.SetVariable(arrayVarName, array.([]int))
	case "[]string":
		array = append(array.([]string), value.(string))
		variable.SetVariable(arrayVarName, array.([]string))
	}
	finished <- true
	return -1
}

// RemoveAt Remove a value by the index, of an array
func RemoveAt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Remove value from array at index ...")

	var array interface{}
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		array = val
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

	switch variable.GetVariableType(array) {
	case "[]bool":
		copy(array.([]bool)[index:], array.([]bool)[index+1:])
		array = array.([]bool)[:len(array.([]bool))-1]
		variable.SetVariable(arrayVarName, array.([]bool))
	case "[]int":
		copy(array.([]int)[index:], array.([]int)[index+1:])
		array = array.([]int)[:len(array.([]int))-1]
		variable.SetVariable(arrayVarName, array.([]int))
	case "[]string":
		copy(array.([]string)[index:], array.([]string)[index+1:])
		array = array.([]string)[:len(array.([]string))-1]
		variable.SetVariable(arrayVarName, array.([]string))
	}
	finished <- true
	return -1
}

// RemoveLast Remove the last value of an array
func RemoveLast(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Remove value from array at end ...")

	var array interface{}
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		array = val
	} else {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		array = array.([]bool)[:len(array.([]bool))-1]
		variable.SetVariable(arrayVarName, array.([]bool))
	case "[]int":
		array = array.([]int)[:len(array.([]int))-1]
		variable.SetVariable(arrayVarName, array.([]int))
	case "[]string":
		array = array.([]string)[:len(array.([]string))-1]
		variable.SetVariable(arrayVarName, array.([]string))
	}
	finished <- true
	return -1
}

// Shuffle an array
func Shuffle(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Shuffling array ...")

	var array interface{}
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		array = val
	} else {
		finished <- true
		return -1
	}

	rand.Seed(time.Now().UnixNano())

	switch variable.GetVariableType(array) {
	case "[]bool":
		rand.Shuffle(len(array.([]bool)), func(i, j int) { array.([]bool)[i], array.([]bool)[j] = array.([]bool)[j], array.([]bool)[i] })
		variable.SetVariable(arrayVarName, array.([]bool))
	case "[]int":
		rand.Shuffle(len(array.([]int)), func(i, j int) { array.([]int)[i], array.([]int)[j] = array.([]int)[j], array.([]int)[i] })
		variable.SetVariable(arrayVarName, array.([]int))
	case "[]string":
		rand.Shuffle(len(array.([]string)), func(i, j int) { array.([]string)[i], array.([]string)[j] = array.([]string)[j], array.([]string)[i] })
		variable.SetVariable(arrayVarName, array.([]string))
	}
	finished <- true
	return -1
}

// UpdateValue Update a value of an array by index
func UpdateValue(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Updating value in array by index ...")

	var array interface{}
	arrayVarName := instructionData.FieldByName("ArrayVarName").Interface().(string)
	if val := variable.SearchVariable(arrayVarName).Value; val != nil {
		array = val
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

	var value interface{}
	valueVarName := instructionData.FieldByName("ValueVarName").Interface().(string)
	if val := variable.SearchVariable(valueVarName).Value; val != nil {
		value = val
	} else {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		array.([]bool)[index] = value.(bool)
		variable.SetVariable(arrayVarName, array.([]bool))
	case "[]int":
		array.([]int)[index] = value.(int)
		variable.SetVariable(arrayVarName, array.([]int))
	case "[]string":
		array.([]string)[index] = value.(string)
		variable.SetVariable(arrayVarName, array.([]string))
	}
	finished <- true
	return -1
}
