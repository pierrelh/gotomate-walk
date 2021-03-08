package sleep

import (
	"fmt"
	"reflect"
)

// Processing process the functions from sleep's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "Sleep":
		duration := instructionData.FieldByName("Duration").Interface().(float64)
		go Sleep(duration, finished)
		<-finished
	case "MilliSleep":
		duration := instructionData.FieldByName("Duration").Interface().(float64)
		go MilliSleep(duration, finished)
		<-finished
	default:
		fmt.Println("FIBER: This function is not integrated yet: " + funcName)
	}
}
