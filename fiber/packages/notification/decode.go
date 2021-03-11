package notification

import "fmt"

// Decode Return the right databinder to decode a saved notification instruction
func Decode(function string) interface{} {
	switch function {
	case "Create":
		return new(CreateDatabinder)
	}
	fmt.Println("ERROR: Unable to find the function")
	return nil
}
