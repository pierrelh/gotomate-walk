package screen

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"

	"github.com/go-vgo/robotgo"
)

// GetMouseColor Get a pixel color by mouse position
func GetMouseColor(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting mouse pixel color ...")
	color := robotgo.GetMouseColor()
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), color)
	finished <- true
	return -1
}

// GetPixelColor Get a pixel color by a position
func GetPixelColor(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting pixel color ...")

	var x int
	if xIsVar := instructionData.FieldByName("XIsVar").Interface().(bool); xIsVar {
		xVarName := instructionData.FieldByName("XVarName").Interface().(string)
		if val := variable.SearchVariable(xVarName).Value; val != nil {
			x = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		x = instructionData.FieldByName("X").Interface().(int)
	}

	var y int
	if yIsVar := instructionData.FieldByName("YIsVar").Interface().(bool); yIsVar {
		yVarName := instructionData.FieldByName("YVarName").Interface().(string)
		if val := variable.SearchVariable(yVarName).Value; val != nil {
			y = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		y = instructionData.FieldByName("Y").Interface().(int)
	}

	color := robotgo.GetPixelColor(x, y)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), color)
	finished <- true
	return -1
}

// GetScreenSize Get the screen size
func GetScreenSize(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting screen size ...")
	w, h := robotgo.GetScreenSize()
	variable.SetVariable(instructionData.FieldByName("HeightOutput").Interface().(string), h)
	variable.SetVariable(instructionData.FieldByName("WidthOutput").Interface().(string), w)
	finished <- true
	return -1
}

func PartScreenShot(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Saving screen shot ...")

	var path string
	if pathIsVar := instructionData.FieldByName("PathIsVar").Interface().(bool); pathIsVar {
		pathVarName := instructionData.FieldByName("PathVarName").Interface().(string)
		if val := variable.SearchVariable(pathVarName).Value; val != nil {
			path = val.(string)
		} else {
			finished <- true
			return -1
		}
	} else {
		path = instructionData.FieldByName("Path").Interface().(string)
	}

	var x int
	if xIsVar := instructionData.FieldByName("XIsVar").Interface().(bool); xIsVar {
		xVarName := instructionData.FieldByName("XVarName").Interface().(string)
		if val := variable.SearchVariable(xVarName).Value; val != nil {
			x = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		x = instructionData.FieldByName("X").Interface().(int)
	}

	var y int
	if yIsVar := instructionData.FieldByName("YIsVar").Interface().(bool); yIsVar {
		yVarName := instructionData.FieldByName("YVarName").Interface().(string)
		if val := variable.SearchVariable(yVarName).Value; val != nil {
			y = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		y = instructionData.FieldByName("Y").Interface().(int)
	}

	var w int
	if wIsVar := instructionData.FieldByName("WIsVar").Interface().(bool); wIsVar {
		wVarName := instructionData.FieldByName("WVarName").Interface().(string)
		if val := variable.SearchVariable(wVarName).Value; val != nil {
			w = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		w = instructionData.FieldByName("W").Interface().(int)
	}

	var h int
	if hIsVar := instructionData.FieldByName("HIsVar").Interface().(bool); hIsVar {
		hVarName := instructionData.FieldByName("HVarName").Interface().(string)
		if val := variable.SearchVariable(hVarName).Value; val != nil {
			h = val.(int)
		} else {
			finished <- true
			return -1
		}
	} else {
		h = instructionData.FieldByName("H").Interface().(int)
	}

	robotgo.SaveCapture(path, x, y, w, h)
	finished <- true
	return -1
}

// SaveCapture Save a screen shot
func ScreenShot(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Saving screen shot ...")

	var path string
	if pathIsVar := instructionData.FieldByName("PathIsVar").Interface().(bool); pathIsVar {
		pathVarName := instructionData.FieldByName("PathVarName").Interface().(string)
		if val := variable.SearchVariable(pathVarName).Value; val != nil {
			path = val.(string)
		} else {
			finished <- true
			return -1
		}
	} else {
		path = instructionData.FieldByName("Path").Interface().(string)
	}

	robotgo.SaveCapture(path)
	finished <- true
	return -1
}
