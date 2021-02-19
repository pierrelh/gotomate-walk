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
func Move(x, y int) {
	robotgo.MoveMouse(x, y)
}

// Scroll scroll the mouse
func Scroll(x, y int) {
	robotgo.Scroll(x, y)
}
