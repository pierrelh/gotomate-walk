package input

import "fmt"

// Decode Return the right databinder to decode a saved input instruction
func Decode(function string) interface{} {
	switch function {
	case "String":
		return new(StringDatabinder)
	case "Int":
		return new(IntDatabinder)
	case "Bool":
		return new(BoolDatabinder)
	}
	fmt.Println("ERROR: Unable to find the function")
	return nil
}
