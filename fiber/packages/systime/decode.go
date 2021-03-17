package systime

import "fmt"

// Decode Return the right databinder to decode a saved systime instruction
func Decode(function string) interface{} {
	switch function {
	case "GetCurrentSysClock":
		return new(ClockDatabinder)
	case "GetCurrentSysTime":
		return new(TimeDatabinder)
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction decode")
	return nil
}
