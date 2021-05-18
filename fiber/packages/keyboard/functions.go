package keyboard

import (
	"fmt"
	"gotomate/fiber/variable"
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

// Type Simulate a type on the keyboard
func Type(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Keyboard type ...")

	var input string
	if isVar := instructionData.FieldByName("InputIsVar").Interface().(bool); isVar {
		varName := instructionData.FieldByName("VarName").Interface().(string)
		if val := variable.SearchVariable(varName).Value; val != nil {
			input = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", varName)
			finished <- true
			return -1
		}
	} else {
		input = instructionData.FieldByName("Input").Interface().(string)
	}

	robotgo.TypeStr(input)
	finished <- true
	return -1
}
