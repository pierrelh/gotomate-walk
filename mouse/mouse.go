package mouse

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// Click Simulate a user click
func Click(btn string, finished chan bool) {
	fmt.Println("Click initialization: ", btn, " click")
	robotgo.Click(btn)
	finished <- true
}

// Move move the mouse
func Move(x, y int, finished chan bool) {
	fmt.Println("Move initialization")
	robotgo.MoveMouse(x, y)
	finished <- true
}

// Scroll scroll the mouse
func Scroll(x, y int, finished chan bool) {
	fmt.Println("Scroll initialization")
	robotgo.Scroll(x, y)
	finished <- true
}
