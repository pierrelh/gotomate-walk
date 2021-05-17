package process

import (
	"fmt"
	"reflect"
)

// Processing process the functions from process's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "KillProcess":
		go func() {
			nextID = KillProcess(instructionData, finished)
		}()
		<-finished
	case "StartProcess":
		go func() {
			nextID = StartProcess(instructionData, finished)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}
