package notification

import (
	"fmt"
	"reflect"
)

// Processing process the functions from notification's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "Create":
		title := instructionData.FieldByName("Title").Interface().(string)
		msg := instructionData.FieldByName("Message").Interface().(string)
		go Create(title, msg, finished)
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
}
