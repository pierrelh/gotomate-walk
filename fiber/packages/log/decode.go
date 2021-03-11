package log

import "fmt"

// Decode Return the right databinder to decode a saved log instruction
func Decode(function string) interface{} {
	switch function {
	case "Print":
		return new(PrintDatabinder)
	}
	fmt.Println("ERROR: Unable to find the function")
	return nil
}
