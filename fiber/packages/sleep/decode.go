package sleep

import "fmt"

// Decode Return the right databinder to decode a saved sleep instruction
func Decode(function string) interface{} {
	switch function {
	case "Sleep":
		return new(Databinder)
	case "MilliSleep":
		return new(Databinder)
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction decode")
	return nil
}
