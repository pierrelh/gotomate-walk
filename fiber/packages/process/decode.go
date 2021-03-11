package process

import "fmt"

// Decode Return the right databinder to decode a saved process instruction
func Decode(function string) interface{} {
	switch function {
	case "StartProcess":
		return new(StartProcessDatabinder)
	}
	fmt.Println("ERROR: Unable to find the function")
	return nil
}
