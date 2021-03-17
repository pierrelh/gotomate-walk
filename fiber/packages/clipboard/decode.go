package clipboard

import "fmt"

// Decode Return the right databinder to decode a saved clipboard instruction
func Decode(function string) interface{} {
	switch function {
	case "Write":
		return new(WriteDatabinder)
	case "Read":
		return new(ReadDatabinder)
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction decode")
	return nil
}
