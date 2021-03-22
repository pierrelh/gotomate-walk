package algorithmic

import (
	"fmt"
	"gotomate/fiber/value"
)

// Define a value in a flow
func Define(name string, values interface{}, finished chan bool) {
	fmt.Println("FIBER INFO: Defining a value ...")
	value.SetValue(name, values)
	finished <- true
}

// If Compare if a statement is true
// func If(valueOne, valueTwo interface{}, comparator string, finished chan bool) bool {
// 	switch comparator {
// 	case "==":
// 		if valueOne == valueTwo {
// 			return true
// 		}
// 		return false
// 	case "!=":
// 		if valueOne != valueTwo {
// 			return true
// 		}
// 		return false
// 	case ">":
// 		if valueOne.(int) > valueTwo.(int) {
// 			return true
// 		}
// 		return false
// 	case ">=":
// 		if valueOne.(int) >= valueTwo.(int) {
// 			return true
// 		}
// 		return false
// 	case "<":
// 		if valueOne.(int) < valueTwo.(int) {
// 			return true
// 		}
// 		return false
// 	case "<=":
// 		if valueOne.(int) <= valueTwo.(int) {
// 			return true
// 		}
// 		return false
// 	}
// 	fmt.Println("| Fiber Finished |")
// 	finished <- true
// 	return false
// }
