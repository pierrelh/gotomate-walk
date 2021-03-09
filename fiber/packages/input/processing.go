package input

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
)

// Processing process the functions from input's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "String":
		go func() {
			msg := instructionData.FieldByName("Message").Interface().(string)
			value.SetValue(instructionData.FieldByName("Output").Interface().(string), String(finished, msg))
		}()
		<-finished
	case "Int":
		go func() {
			msg := instructionData.FieldByName("Message").Interface().(string)
			value.SetValue(instructionData.FieldByName("Output").Interface().(string), Int(finished, msg))
		}()
		<-finished
	case "Bool":
		go func() {
			msg := instructionData.FieldByName("Message").Interface().(string)
			value.SetValue(instructionData.FieldByName("Output").Interface().(string), Bool(finished, msg))
		}()
		<-finished
	default:
		fmt.Println("FIBER: This function is not integrated yet: " + funcName)
	}
}
