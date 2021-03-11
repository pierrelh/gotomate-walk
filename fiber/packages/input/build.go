package input

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a input instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "String":
		return new(StringDatabinder), InputTemplate
	case "Int":
		return new(IntDatabinder), InputTemplate
	case "Bool":
		return new(BoolDatabinder), InputTemplate
	}
	fmt.Println("ERROR: Unable to find the function")
	return nil, nil
}
