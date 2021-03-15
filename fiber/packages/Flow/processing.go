package flow

import (
	"fmt"
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "Start":
		go Start(finished)
		<-finished
	case "End":
		go End(finished)
		<-finished
	default:
		fmt.Println("FIBER: This function is not integrated yet: " + funcName)
	}
}
