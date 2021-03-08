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
			newValue := &value.InstructionValue{
				Key:   instructionData.FieldByName("Output").Interface().(string),
				Value: StartProcess(finished, path),
			}
			value.FiberValues = append(value.FiberValues, newValue)
		}()
		<-finished
	default:
		fmt.Println("FIBER: This function is not integrated yet: " + funcName)
	}
}
