package screen

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
	"strconv"

	"github.com/go-vgo/robotgo"
)

// GetMouseColor Get a pixel color by mouse position
func GetMouseColor(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting mouse pixel color ...")
	color := robotgo.GetMouseColor()
	value.SetValue(instructionData.FieldByName("Output").Interface().(string), color)
	finished <- true
	return -1
}

// GetPixelColor Get a pixel color by a position
func GetPixelColor(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting pixel color ...")
	Sx := instructionData.FieldByName("X").Interface().(string)
	x, _ := strconv.Atoi(Sx)
	if xIsVar := instructionData.FieldByName("XIsVar").Interface().(bool); xIsVar {
		if val := value.KeySearch(Sx).Value; val != nil {
			x = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", Sx)
			finished <- true
			return -1
		}
	}

	Sy := instructionData.FieldByName("Y").Interface().(string)
	y, _ := strconv.Atoi(Sy)
	if yIsVar := instructionData.FieldByName("YIsVar").Interface().(bool); yIsVar {
		if val := value.KeySearch(Sy).Value; val != nil {
			y = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", Sy)
			finished <- true
			return -1
		}
	}

	color := robotgo.GetPixelColor(x, y)
	value.SetValue(instructionData.FieldByName("Output").Interface().(string), color)
	finished <- true
	return -1
}

// GetScreenSize Get the screen size
func GetScreenSize(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting screen size ...")
	w, h := robotgo.GetScreenSize()
	value.SetValue(instructionData.FieldByName("HeightOutput").Interface().(string), h)
	value.SetValue(instructionData.FieldByName("WidthOutput").Interface().(string), w)
	finished <- true
	return -1
}

// SaveCapture Save a screen shot
func SaveCapture(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Saving screen shot ...")
	path := instructionData.FieldByName("Path").Interface().(string)
	if pathIsVar := instructionData.FieldByName("PathIsVar").Interface().(bool); pathIsVar {
		if val := value.KeySearch(path).Value; val != nil {
			path = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", path)
			finished <- true
			return -1
		}
	}

	robotgo.SaveCapture(path)
	finished <- true
	return -1
}
