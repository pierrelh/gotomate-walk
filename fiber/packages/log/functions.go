package log

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"
)

// Print log a value
func Print(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Logging ...")

	var log interface{}
	if isVar := instructionData.FieldByName("LogIsVar").Interface().(bool); isVar {
		varName := instructionData.FieldByName("VarName").Interface().(string)
		if val := variable.SearchVariable(varName).Value; val != nil {
			log = val
		} else {
			finished <- true
			return -1
		}
	} else {
		log = instructionData.FieldByName("Log").Interface()
	}
	fmt.Println("LOG: ", log)
	finished <- true
	return -1
}
