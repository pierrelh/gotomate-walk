package mouse

import "fmt"

// Decode Return the right databinder to decode a saved mouse instruction
func Decode(function string) interface{} {
	switch function {
	case "Click":
		return new(ClickDatabinder)
	case "Scroll":
		return new(ScrollDatabinder)
	case "Move":
		return new(MoveDatabinder)
	}
	fmt.Println("ERROR: Unable to find the function")
	return nil
}
