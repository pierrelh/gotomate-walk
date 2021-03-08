package screen

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
)

// Processing process the functions from screen's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	case "GetMouseColor":
		go func() {
			newValue := &value.InstructionValue{
				Key:   instructionData.FieldByName("Output").Interface().(string),
				Value: GetMouseColor(finished),
			}
			value.FiberValues = append(value.FiberValues, newValue)
		}()
		<-finished
	case "GetPixelColor":
		go func() {
			x := instructionData.FieldByName("X").Interface().(int)
			y := instructionData.FieldByName("Y").Interface().(int)
			newValue := &value.InstructionValue{
				Key:   instructionData.FieldByName("Output").Interface().(string),
				Value: GetPixelColor(x, y, finished),
			}
			value.FiberValues = append(value.FiberValues, newValue)
		}()
		<-finished
	case "SaveCapture":
		path := instructionData.FieldByName("Path").Interface().(string)
		go SaveCapture(finished, path)
		<-finished
	default:
		fmt.Println("FIBER: This function is not integrated yet: " + funcName)
	}
}
