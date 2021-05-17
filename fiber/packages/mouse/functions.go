package mouse

import (
	"fmt"
	"gotomate/fiber/value"
	"reflect"
	"strconv"

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

// Move move the mouse
func Move(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Moving mouse ...")
	xs := instructionData.FieldByName("X").Interface().(string)
	x, _ := strconv.Atoi(xs)
	if xIsVar := instructionData.FieldByName("XIsVar").Interface().(bool); xIsVar {
		if val := value.KeySearch(xs).Value; val != nil {
			x = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", xs)
			finished <- true
			return -1
		}
	}

	ys := instructionData.FieldByName("Y").Interface().(string)
	y, _ := strconv.Atoi(ys)
	if yIsVar := instructionData.FieldByName("YIsVar").Interface().(bool); yIsVar {
		if val := value.KeySearch(ys).Value; val != nil {
			y = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", ys)
			finished <- true
			return -1
		}
	}

	robotgo.MoveMouse(x, y)
	finished <- true
	return -1
}

// Scroll scroll the mouse
func Scroll(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Scrolling mouse ...")
	xs := instructionData.FieldByName("X").Interface().(string)
	x, _ := strconv.Atoi(xs)
	if xIsVar := instructionData.FieldByName("XIsVar").Interface().(bool); xIsVar {
		if val := value.KeySearch(xs).Value; val != nil {
			x = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", xs)
			finished <- true
			return -1
		}
	}

	ys := instructionData.FieldByName("Y").Interface().(string)
	y, _ := strconv.Atoi(ys)
	if yIsVar := instructionData.FieldByName("YIsVar").Interface().(bool); yIsVar {
		if val := value.KeySearch(ys).Value; val != nil {
			y = val.(int)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", ys)
			finished <- true
			return -1
		}
	}

	robotgo.Scroll(x, y)
	finished <- true
	return -1
}
