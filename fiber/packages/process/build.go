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
	}
	fmt.Println("ERROR: Unable to find the function")
	return nil, nil
}
