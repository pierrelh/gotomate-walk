package mouse

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a mouse instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Click":
		return new(ClickDatabinder), ClickTemplate
	case "Scroll":
		return new(ScrollDatabinder), ScrollTemplate
	case "Move":
		return new(MoveDatabinder), MoveTemplate
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}
