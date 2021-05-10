package algorithmic

import (
	"fmt"
	"gotomate/fiber/value"
)

// DefineInt an int value in a flow
func DefineInt(name string, val int, finished chan bool) {
	fmt.Println("FIBER INFO: Defining an int value ...")
	value.SetValue(name, val)
	finished <- true
}

// DefineString a string value in a flow
func DefineString(name string, val string, finished chan bool) {
	fmt.Println("FIBER INFO: Defining a string value ...")
	value.SetValue(name, val)
	finished <- true
}

// DefineBool a bool value in a flow
func DefineBool(name string, val bool, finished chan bool) {
	fmt.Println("FIBER INFO: Defining a bool value ...")
	value.SetValue(name, val)
	finished <- true
}

// If Compare if a statement is true
func If(valueOne, valueTwo interface{}, comparator string, finished chan bool) bool {
	fmt.Println("FIBER INFO: If Statement ...")
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
	return statement
}
