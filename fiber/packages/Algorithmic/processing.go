package algorithmic

import (
	"fmt"
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "DefineBool":
		go func() {
			nextID = DefineBool(instructionData, finished)
		}()
		<-finished
	case "DefineInt":
		go func() {
			nextID = DefineInt(instructionData, finished)
		}()
		<-finished
	case "DefineString":
		go func() {
			nextID = DefineString(instructionData, finished)
		}()
		<-finished
	case "If":
		go func() {
			nextID = If(instructionData, finished)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}
