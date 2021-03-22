package algorithmic

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Define":
		return new(DefineDatabinder), DefineTemplate
		// case "If":
		// 	return new(IfDatabinder), IfTemplate
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}
