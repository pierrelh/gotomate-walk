package process

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
)

// Processing process the functions from process's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "StartProcess":
		go func() {
			path := instructionData.FieldByName("Path").Interface().(string)
			value.SetValue(instructionData.FieldByName("Output").Interface().(string), StartProcess(finished, path))
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
}
