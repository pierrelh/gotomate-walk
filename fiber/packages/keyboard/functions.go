package keyboard

import (
	"fmt"
	"reflect"

	"github.com/go-vgo/robotgo"
)

// Tap Simulate a tap on the keyboard
func Tap(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Keyboard tap ...")
	input := instructionData.FieldByName("Input").Interface().(string)
	special := []string{}
	if err := instructionData.FieldByName("Special1").Interface().(string); err != "" {
		special = append(special, err)
	}
	if err := instructionData.FieldByName("Special2").Interface().(string); err != "" {
		special = append(special, err)
	}
	if len(special) == 0 {
		robotgo.KeyTap(input)
	} else {
		robotgo.KeyTap(input, special)
	}
	finished <- true
	return -1
}
