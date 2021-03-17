package screen

import "fmt"

// Decode Return the right databinder to decode a saved screen instruction
func Decode(function string) interface{} {
	switch function {
	case "GetPixelColor":
		return new(PixelColorDatabinder)
	case "GetMouseColor":
		return new(MouseColorDatabinder)
	case "SaveCapture":
		return new(SaveCaptureDatabinder)
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction decode")
	return nil
}
