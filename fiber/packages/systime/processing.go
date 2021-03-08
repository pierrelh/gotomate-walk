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
			clock := [3]int{h, m, s}
			newValue := &value.InstructionValue{
				Key:   instructionData.FieldByName("Output").Interface().(string),
				Value: clock,
			}
			value.FiberValues = append(value.FiberValues, newValue)
		}()
		<-finished
	case "GetCurrentSysTime":
		go func() {
			newValue := &value.InstructionValue{
				Key:   instructionData.FieldByName("Output").Interface().(string),
				Value: GetCurrentSysTime(finished),
			}
			value.FiberValues = append(value.FiberValues, newValue)
		}()
		<-finished
	default:
		fmt.Println("FIBER: This function is not integrated yet: " + funcName)
	}
}
