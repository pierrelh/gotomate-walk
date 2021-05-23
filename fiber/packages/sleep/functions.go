package sleep

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"
	"time"
)

// MilliSleep sleep tm milli second
func MilliSleep(instructionData reflect.Value, finished chan bool) int {
	var duration float64
	if durationIsVar := instructionData.FieldByName("DurationIsVar").Interface().(bool); durationIsVar {
		durationVarName := instructionData.FieldByName("DurationVarName").Interface().(string)
		if val := variable.SearchVariable(durationVarName).Value; val != nil {
			duration = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		duration = instructionData.FieldByName("Duration").Interface().(float64)
	}

	fmt.Println("FIBER INFO: Sleeping for: ", duration, "ms")
	time.Sleep(time.Duration(duration) * time.Millisecond)
	finished <- true
	return -1
}

// Sleep time.Sleep tm second
func Sleep(instructionData reflect.Value, finished chan bool) int {
	var duration float64
	if durationIsVar := instructionData.FieldByName("DurationIsVar").Interface().(bool); durationIsVar {
		durationVarName := instructionData.FieldByName("DurationVarName").Interface().(string)
		if val := variable.SearchVariable(durationVarName).Value; val != nil {
			duration = val.(float64)
		} else {
			finished <- true
			return -1
		}
	} else {
		duration = instructionData.FieldByName("Duration").Interface().(float64)
	}

	fmt.Println("FIBER INFO: Sleeping for: ", duration, "s")
	time.Sleep(time.Duration(duration) * time.Second)
	finished <- true
	return -1
}
