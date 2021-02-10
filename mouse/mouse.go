package mouse

/*
#include "mouse_functions/goMouse.h"
*/
import "C"

import "unsafe"

/*
.___  ___.   ______    __    __       _______. _______
|   \/   |  /  __  \  |  |  |  |     /       ||   ____|
|  \  /  | |  |  |  | |  |  |  |    |   (----`|  |__
|  |\/|  | |  |  |  | |  |  |  |     \   \    |   __|
|  |  |  | |  `--'  | |  `--'  | .----)   |   |  |____
|__|  |__|  \______/   \______/  |_______/    |_______|

*/

// CheckMouse check the mouse button
func CheckMouse(btn string) C.MMMouseButton {
	// button = args[0].(C.MMMouseButton)
	if btn == "left" {
		return C.LEFT_BUTTON
	}

	if btn == "center" {
		return C.CENTER_BUTTON
	}

	if btn == "right" {
		return C.RIGHT_BUTTON
	}

	return C.LEFT_BUTTON
}

// MoveMouse move the mouse
func MoveMouse(x, y int) {
	// C.size_t  int
	Move(x, y)
}

// Move move the mouse
func Move(x, y int) {
	cx := C.int32_t(x)
	cy := C.int32_t(y)
	C.move_mouse(cx, cy)
}

// DragMouse drag the mouse
func DragMouse(x, y int, args ...string) {
	Drag(x, y, args...)
}

// Drag drag the mouse
func Drag(x, y int, args ...string) {
	var button C.MMMouseButton = C.LEFT_BUTTON

	cx := C.int32_t(x)
	cy := C.int32_t(y)

	if len(args) > 0 {
		button = CheckMouse(args[0])
	}

	C.drag_mouse(cx, cy, button)
}

// DragSmooth drag the mouse smooth
func DragSmooth(x, y int, args ...interface{}) {
	MouseToggle("down")
	MoveSmooth(x, y, args...)
	MouseToggle("up")
}

// MoveMouseSmooth move the mouse smooth,
// moves mouse to x, y human like, with the mouse button up.
func MoveMouseSmooth(x, y int, args ...interface{}) bool {
	return MoveSmooth(x, y, args...)
}

// MoveSmooth move the mouse smooth,
// moves mouse to x, y human like, with the mouse button up.
//
// robotgo.MoveSmooth(x, y int, low, high float64, mouseDelay int)
func MoveSmooth(x, y int, args ...interface{}) bool {
	cx := C.int32_t(x)
	cy := C.int32_t(y)

	var (
		mouseDelay = 10
		low        C.double
		high       C.double
	)

	if len(args) > 2 {
		mouseDelay = args[2].(int)
	}

	if len(args) > 1 {
		low = C.double(args[0].(float64))
		high = C.double(args[1].(float64))
	} else {
		low = 1.0
		high = 3.0
	}

	cbool := C.move_mouse_smooth(cx, cy, low, high, C.int(mouseDelay))

	return bool(cbool)
}

// MoveArgs move mose relative args
func MoveArgs(x, y int) (int, int) {
	mx, my := GetMousePos()
	mx = mx + x
	my = my + y

	return mx, my
}

// MoveRelative move mose relative
func MoveRelative(x, y int) {
	Move(MoveArgs(x, y))
}

// MoveSmoothRelative move mose smooth relative
func MoveSmoothRelative(x, y int, args ...interface{}) {
	mx, my := MoveArgs(x, y)
	MoveSmooth(mx, my, args...)
}

// GetMousePos get mouse's portion
func GetMousePos() (int, int) {
	pos := C.get_mouse_pos()

	x := int(pos.x)
	y := int(pos.y)

	return x, y
}

// MouseClick click the mouse
//
// robotgo.MouseClick(button string, double bool)
func MouseClick(args ...interface{}) {
	Click(args...)
}

// Click click the mouse
//
// robotgo.Click(button string, double bool)
func Click(args ...interface{}) {
	var (
		button C.MMMouseButton = C.LEFT_BUTTON
		double C.bool
	)

	if len(args) > 0 {
		button = CheckMouse(args[0].(string))
	}

	if len(args) > 1 {
		double = C.bool(args[1].(bool))
	}

	C.mouse_click(button, double)
}

// MoveClick move and click the mouse
//
// robotgo.MoveClick(x, y int, button string, double bool)
func MoveClick(x, y int, args ...interface{}) {
	MoveMouse(x, y)
	MouseClick(args...)
}

// MovesClick move smooth and click the mouse
func MovesClick(x, y int, args ...interface{}) {
	MoveSmooth(x, y)
	MouseClick(args...)
}

// MouseToggle toggle the mouse
func MouseToggle(togKey string, args ...interface{}) int {
	var button C.MMMouseButton = C.LEFT_BUTTON

	if len(args) > 0 {
		button = CheckMouse(args[0].(string))
	}

	down := C.CString(togKey)
	i := C.mouse_toggle(down, button)

	C.free(unsafe.Pointer(down))
	return int(i)
}

// ScrollMouse scroll the mouse
func ScrollMouse(x int, direction string) {
	cx := C.size_t(x)
	cy := C.CString(direction)
	C.scroll_mouse(cx, cy)

	C.free(unsafe.Pointer(cy))
}

// Scroll scroll the mouse with x, y
//
// robotgo.Scroll(x, y, msDelay int)
func Scroll(x, y int, args ...int) {
	var msDelay = 10
	if len(args) > 0 {
		msDelay = args[0]
	}

	cx := C.int(x)
	cy := C.int(y)
	cz := C.int(msDelay)

	C.scroll(cx, cy, cz)
}

// SetMouseDelay set mouse delay
func SetMouseDelay(delay int) {
	cdelay := C.size_t(delay)
	C.set_mouse_delay(cdelay)
}
