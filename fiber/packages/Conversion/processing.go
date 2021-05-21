package conversion

import (
	"fmt"
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "IntToString":
		go func() {
			nextID = IntToString(instructionData, finished)
		}()
		<-finished
	case "StringToBool":
		go func() {
			nextID = StringToBool(instructionData, finished)
		}()
		<-finished
	case "StringToInt":
		go func() {
			nextID = StringToInt(instructionData, finished)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}
