package systime

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
	"time"
)

// GetCurrentSysClock return the current sys clock
func GetCurrentSysClock(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting sys clock...")
	h, m, s := time.Now().Clock()
	value.SetValue(instructionData.FieldByName("HoursOutput").Interface().(string), h)
	value.SetValue(instructionData.FieldByName("MinutesOutput").Interface().(string), m)
	value.SetValue(instructionData.FieldByName("SecondsOutput").Interface().(string), s)
	finished <- true
	return -1
}

// GetCurrentSysTime return the current sys time
func GetCurrentSysTime(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting sys time...")
	t := time.Now()
	value.SetValue(instructionData.FieldByName("Output").Interface().(string), t)
	finished <- true
	return -1
}
