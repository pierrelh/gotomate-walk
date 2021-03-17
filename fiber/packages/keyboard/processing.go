package keyboard

import (
	"fmt"
	"reflect"
)

// Processing process the functions from keyboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "Tap":
		input := instructionData.FieldByName("Input").Interface().(string)
		special := []string{}
		if err := instructionData.FieldByName("Special1").Interface().(string); err != "" {
			special = append(special, err)
		}
		if err := instructionData.FieldByName("Special2").Interface().(string); err != "" {
			special = append(special, err)
		}
		go Tap(input, special, finished)
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
}
