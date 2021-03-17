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
			value.SetValue(instructionData.FieldByName("Output").Interface().(string), content)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
}
