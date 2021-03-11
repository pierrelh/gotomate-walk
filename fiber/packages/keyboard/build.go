package keyboard

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a keyboard instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Tap":
		return new(TapDatabinder), TapTemplate
	}
	fmt.Println("ERROR: Unable to find the function")
	return nil, nil
}
