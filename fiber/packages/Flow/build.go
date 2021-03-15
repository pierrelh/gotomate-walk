package flow

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Start":
		return new(StartDatabinder), StartTemplate
	case "End":
		return new(EndDatabinder), nil
	}
	fmt.Println("ERROR: Unable to find the function")
	return nil, nil
}
