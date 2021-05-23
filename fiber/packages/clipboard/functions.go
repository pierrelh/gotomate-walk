package clipboard

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"

	"github.com/go-vgo/robotgo"
)

// Read read string from clipboard
func Read(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Clipboard Reading ...")
	content, _ := robotgo.ReadAll()
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), content)
	finished <- true
	return -1
}

// Write write string to clipboard
func Write(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Clipboard Writing ...")

	var content string
	if isVar := instructionData.FieldByName("IsVar").Interface().(bool); isVar {
		varName := instructionData.FieldByName("VarName").Interface().(string)
		if val := variable.SearchVariable(varName).Value; val != nil {
			content = val.(string)
		} else {
			finished <- true
			return -1
		}
	} else {
		content = instructionData.FieldByName("Content").Interface().(string)
	}
	robotgo.WriteAll(content)
	finished <- true
	return -1
}
