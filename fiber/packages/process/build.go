package process

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a process instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "StartProcess":
		return new(StartProcessDatabinder), StartProcessTemplate
	case "KillProcess":
		return new(KillProcessDatabinder), KillProcessTemplate
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}
