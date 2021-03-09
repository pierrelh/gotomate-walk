package keyboard

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// Tap Simulate a tap on the keyboard
func Tap(key string, sp []string, finished chan bool) {
	fmt.Println("FIBER: Keyboard tap ...")
	if len(sp) == 0 {
		robotgo.KeyTap(key)
	} else {
		robotgo.KeyTap(key, sp)
	}
	finished <- true
}
