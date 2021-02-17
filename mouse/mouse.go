package mouse

import "github.com/go-vgo/robotgo"

// Click Simulate a user click
func Click() {
	robotgo.MouseClick()
}

// Move move the mouse
func Move(x, y int) {
	robotgo.MoveMouse(x, y)
}

// Scroll scroll the mouse
func Scroll(x, y int) {
	robotgo.Scroll(x, y)
}
