package mouse

import (
	"fmt"
	"reflect"
)

// Processing process the functions from mouse's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "Click":
		button := instructionData.FieldByName("MouseButtonName").Interface().(string)
		go Click(button, finished)
		<-finished
	case "Scroll":
		x := instructionData.FieldByName("X").Interface().(int)
		y := instructionData.FieldByName("Y").Interface().(int)
		go Scroll(x, y, finished)
		<-finished
	case "Move":
		x := instructionData.FieldByName("X").Interface().(int)
		y := instructionData.FieldByName("Y").Interface().(int)
		go Move(x, y, finished)
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
}
