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
	fmt.Println("ERROR: Unable to find the function")
	return nil
}
