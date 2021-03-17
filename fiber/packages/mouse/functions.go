package mouse

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// Click Simulate a user click
func Click(btn string, finished chan bool) {
	fmt.Println("FIBER INFO: Clicking mouse  ...")
	robotgo.Click(btn)
	finished <- true
}

// Move move the mouse
func Move(x, y int, finished chan bool) {
	fmt.Println("FIBER INFO: Moving mouse ...")
	robotgo.MoveMouse(x, y)
	finished <- true
}

// Scroll scroll the mouse
func Scroll(x, y int, finished chan bool) {
	fmt.Println("FIBER INFO: Scrolling mouse ...")
	robotgo.Scroll(x, y)
	finished <- true
}
