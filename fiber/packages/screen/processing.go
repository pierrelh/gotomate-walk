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
			value.SetValue(instructionData.FieldByName("Output").Interface().(string), GetMouseColor(finished))
		}()
		<-finished
	case "GetPixelColor":
		go func() {
			x := instructionData.FieldByName("X").Interface().(int)
			y := instructionData.FieldByName("Y").Interface().(int)
			value.SetValue(instructionData.FieldByName("Output").Interface().(string), GetPixelColor(x, y, finished))
		}()
		<-finished
	case "GetScreenSize":
		go func() {
			w, h := GetScreenSize(finished)
			value.SetValue(instructionData.FieldByName("HeightOutput").Interface().(string), h)
			value.SetValue(instructionData.FieldByName("WidthOutput").Interface().(string), w)
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
