package mouse

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"

	"github.com/go-vgo/robotgo"
)

// Click Simulate a user click
func Click(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Clicking mouse  ...")
	button := instructionData.FieldByName("MouseButtonName").Interface().(string)
	robotgo.Click(button)
	finished <- true
	return -1
}

// Drag Simulate a user drag
func Drag(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Clicking mouse  ...")

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

	robotgo.MouseToggle("down")
	robotgo.Drag(x, y)
	robotgo.MouseToggle("up")
	finished <- true
	return -1
}

// DragSmooth Simulate a user draggin smoothly
func DragSmooth(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Clicking mouse  ...")

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

	robotgo.MouseToggle("down")
	robotgo.DragSmooth(x, y)
	robotgo.MouseToggle("up")
	finished <- true
	return -1
}

// Move move the mouse
func Move(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Moving mouse ...")

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

	robotgo.MoveMouse(x, y)
	finished <- true
	return -1
}

// MoveSmooth move the mouse smoothly
func MoveSmooth(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Moving mouse ...")

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

	robotgo.MoveSmooth(x, y)
	finished <- true
	return -1
}

// Scroll scroll the mouse
func Scroll(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Scrolling mouse ...")

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

	robotgo.Scroll(x, y)
	finished <- true
	return -1
}
