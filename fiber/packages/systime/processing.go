package systime

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
)

// Processing process the functions from systime's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "GetCurrentSysClock":
		go func() {
			h, m, s := GetCurrentSysClock(finished)
			value.SetValue(instructionData.FieldByName("HoursOutput").Interface().(string), h)
			value.SetValue(instructionData.FieldByName("MinutesOutput").Interface().(string), m)
			value.SetValue(instructionData.FieldByName("SecondsOutput").Interface().(string), s)
		}()
		<-finished
	case "GetCurrentSysTime":
		go func() {
			value.SetValue(instructionData.FieldByName("Output").Interface().(string), GetCurrentSysTime(finished))
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
}
