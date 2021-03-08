package input

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
)

// Processing process the functions from input's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "String":
		go func() {
			msg := instructionData.FieldByName("Message").Interface().(string)
			newValue := &value.InstructionValue{
				Key:   instructionData.FieldByName("Output").Interface().(string),
				Value: String(finished, msg),
			}
			value.FiberValues = append(value.FiberValues, newValue)
		}()
		<-finished
	case "Int":
		go func() {
			msg := instructionData.FieldByName("Message").Interface().(string)
			newValue := &value.InstructionValue{
				Key:   instructionData.FieldByName("Output").Interface().(string),
				Value: Int(finished, msg),
			}
			value.FiberValues = append(value.FiberValues, newValue)
		}()
		<-finished
	case "Bool":
		go func() {
			msg := instructionData.FieldByName("Message").Interface().(string)
			newValue := &value.InstructionValue{
				Key:   instructionData.FieldByName("Output").Interface().(string),
				Value: Bool(finished, msg),
			}
			value.FiberValues = append(value.FiberValues, newValue)
		}()
		<-finished
	default:
		fmt.Println("FIBER: This function is not integrated yet: " + funcName)
	}
}
