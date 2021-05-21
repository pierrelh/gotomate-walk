package conversion

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a clipboard instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "IntToString":
		return new(IntToStringDatabinder), IntToStringTemplate
	case "StringToBool":
		return new(StringConversionDatabinder), StringToBoolTemplate
	case "StringToInt":
		return new(StringConversionDatabinder), StringToIntTemplate
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}
