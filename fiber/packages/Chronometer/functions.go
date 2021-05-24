package chronometer

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"
	"time"
)

// End a chronometer & save the elapsed time
func End(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting elapsed chronometer time ...")

	var chrono time.Time
	chronoVarName := instructionData.FieldByName("ChronoVarName").Interface().(string)
	if val := variable.SearchVariable(chronoVarName).Value; val != nil {
		chrono = val.(time.Time)
	} else {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), time.Since(chrono))
	finished <- true
	return -1
}

// Start a chronometer
func Start(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting a new chronometer ...")

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), time.Now())
	finished <- true
	return -1
}
