package keyboard

import "fmt"

// Decode Return the right databinder to decode a saved keyboard instruction
func Decode(function string) interface{} {
	switch function {
	case "Tap":
		return new(TapDatabinder)
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction decode")
	return nil
}
