package screen

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// GetPixelColor Get a pixel color by a position
func GetPixelColor(x, y int, finished chan bool) string {
	fmt.Println("FIBER INFO: Getting pixel color ...")
	color := robotgo.GetPixelColor(x, y)
	finished <- true
	return color
}

// GetMouseColor Get a pixel color by mouse position
func GetMouseColor(finished chan bool) string {
	fmt.Println("FIBER INFO: Getting mouse pixel color ...")
	color := robotgo.GetMouseColor()
	finished <- true
	return color
}

// GetScreenSize Get the screen size
func GetScreenSize(finished chan bool) (int, int) {
	fmt.Println("FIBER INFO: Getting screen size ...")
	w, h := robotgo.GetScreenSize()
	finished <- true
	return w, h
}

// SaveCapture Save a screen shot
func SaveCapture(finished chan bool, spath string) {
	fmt.Println("FIBER INFO: Saving screen shot ...")
	robotgo.SaveCapture(spath)
	finished <- true
}
