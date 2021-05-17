package sleep

import (
	"fmt"
	"reflect"
	"time"
)

// MilliSleep sleep tm milli second
func MilliSleep(instructionData reflect.Value, finished chan bool) int {
	duration := instructionData.FieldByName("Duration").Interface().(float64)
	fmt.Println("FIBER INFO: Sleeping for: ", duration, "ms")
	time.Sleep(time.Duration(duration) * time.Millisecond)
	finished <- true
	return -1
}

// Sleep time.Sleep tm second
func Sleep(instructionData reflect.Value, finished chan bool) int {
	duration := instructionData.FieldByName("Duration").Interface().(float64)
	fmt.Println("FIBER INFO: Sleeping for: ", duration, "s")
	time.Sleep(time.Duration(duration) * time.Second)
	finished <- true
	return -1
}
