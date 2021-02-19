package keyboard

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// Tap Simulate a tap on the keyboard
func Tap(key string, sp []string, finished chan bool) {
	fmt.Println("Tap initialization")
	if len(sp) == 0 {
		robotgo.KeyTap(key)
	} else {
		robotgo.KeyTap(key, sp)
	}
	finished <- true
}
