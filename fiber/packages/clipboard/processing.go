package clipboard

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "Write":
		content := instructionData.FieldByName("Content").Interface().(string)
		go Write(content, finished)
		<-finished
	case "Read":
		go func() {
			content, _ := Read(finished)
			newValue := &value.InstructionValue{
				Key:   instructionData.FieldByName("Output").Interface().(string),
				Value: content,
			}
			value.FiberValues = append(value.FiberValues, newValue)
		}()
		<-finished
	default:
		fmt.Println("FIBER: This function is not integrated yet: " + funcName)
	}
}
