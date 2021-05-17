package algorithmic

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
	"strconv"
)

// DefineBool a bool value in a flow
func DefineBool(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Defining a bool value ...")
	name := instructionData.FieldByName("Name").Interface().(string)
	setVal := instructionData.FieldByName("Value").Interface().(string)
	boolVal, _ := strconv.ParseBool(setVal)

	if isVar := instructionData.FieldByName("IsVar").Interface().(bool); isVar {
		if val := value.KeySearch(setVal).Value; val != nil {
			boolVal = val.(bool)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", setVal)
			finished <- true
			return -1
		}
	}
	value.SetValue(name, boolVal)
	finished <- true
	return -1
}

// DefineInt an int value in a flow
func DefineInt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Defining an int value ...")
	name := instructionData.FieldByName("Name").Interface().(string)
	setVal := instructionData.FieldByName("Value").Interface().(string)
	intVal, _ := strconv.Atoi(setVal)

	if isVar := instructionData.FieldByName("IsVar").Interface().(bool); isVar {
		if val := value.KeySearch(setVal).Value; val != nil {
			intVal = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", setVal)
			finished <- true
			return -1
		}
	}
	value.SetValue(name, intVal)
	finished <- true
	return -1
}

// DefineString a string value in a flow
func DefineString(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Defining a string value ...")
	name := instructionData.FieldByName("Name").Interface().(string)
	setVal := instructionData.FieldByName("Value").Interface().(string)

	if isVar := instructionData.FieldByName("IsVar").Interface().(bool); isVar {
		if val := value.KeySearch(setVal).Value; val != nil {
			setVal = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", setVal)
			finished <- true
			return -1
		}
	}
	value.SetValue(name, setVal)
	finished <- true
	return -1
}

// If Compare if a statement is true
func If(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: If Statement ...")
	valueOne := instructionData.FieldByName("ValueOne").Interface()

	if oneIsVar := instructionData.FieldByName("OneIsVar").Interface().(bool); oneIsVar {
		if val := value.KeySearch(valueOne.(string)).Value; val != nil {
			valueOne = val
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", valueOne)
			finished <- true
			return -1
		}
	}
	valueTwo := instructionData.FieldByName("ValueTwo").Interface()

	if twoIsVar := instructionData.FieldByName("TwoIsVar").Interface().(bool); twoIsVar {
		if val := value.KeySearch(valueTwo.(string)).Value; val != nil {
			valueTwo = val
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", valueTwo)
			finished <- true
			return -1
		}
	}
	comparator := instructionData.FieldByName("Comparator").Interface().(string)
	falseInstruction := instructionData.FieldByName("FalseInstruction").Interface().(int)
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
		if valueOne.(int) > valueTwo.(int) {
			statement = true
		}
	case ">=":
		if valueOne.(int) >= valueTwo.(int) {
			statement = true
		}
	case "<":
		if valueOne.(int) < valueTwo.(int) {
			statement = true
		}
	case "<=":
		if valueOne.(int) <= valueTwo.(int) {
			statement = true
		}
	}
	finished <- true
	if statement {
		return -1
	} else {
		return falseInstruction
	}
}
