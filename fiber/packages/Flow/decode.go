package flow

import "fmt"

// Decode Return the right databinder to decode a saved clipboard instruction
func Decode(function string) interface{} {
	switch function {
	case "Start":
		return new(StartDatabinder)
	case "End":
		return new(EndDatabinder)
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction decode")
	return nil
}
