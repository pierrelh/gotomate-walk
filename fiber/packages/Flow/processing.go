package flow

import (
	"fmt"
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) bool {
	switch funcName {
	case "Start":
		go Start(finished)
		<-finished
		return false
	case "End":
		go End(finished)
		<-finished
		return true
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
		return false
	}
}
