package clipboard

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"

	"github.com/go-vgo/robotgo"
)

// Read read string from clipboard
func Read(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Clipboard Reading ...")
	content, _ := robotgo.ReadAll()
	value.SetValue(instructionData.FieldByName("Output").Interface().(string), content)
	finished <- true
	return -1
}

// Write write string to clipboard
func Write(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Clipboard Writing ...")
	content := instructionData.FieldByName("Content").Interface().(string)

	if isVar := instructionData.FieldByName("ContentIsVar").Interface().(bool); isVar {
		if val := value.KeySearch(content).Value; val != nil {
			content = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", content)
			finished <- true
			return -1
		}
	}
	robotgo.WriteAll(content)
	finished <- true
	return -1
}
