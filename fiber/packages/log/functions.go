package log

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
)

// Print log a value
func Print(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Logging ...")
	log := instructionData.FieldByName("Log").Interface()

	if isVar := instructionData.FieldByName("LogIsVar").Interface().(bool); isVar {
		if val := value.KeySearch(log.(string)).Value; val != nil {
			log = val
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", log)
			finished <- true
			return -1
		}
	}
	fmt.Println("LOG: ", log)
	finished <- true
	return -1
}
