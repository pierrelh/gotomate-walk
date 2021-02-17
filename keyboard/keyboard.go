package keyboard

import "github.com/go-vgo/robotgo"

// Tap Simulate a tap on the keyboard
func Tap(key string) {
	robotgo.KeyTap(key)
}
